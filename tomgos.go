package tomgos

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
)

type StructFields struct {
	Name           string
	Type           string
	JSONDescriptor string
}

type Struct struct {
	StructName string
	Fields     []StructFields
}

type GeneratedData struct {
	PackageName string
	Structs     []Struct
	UsingTime   bool
}

type Generator interface {
	Generate(tomlPathFile string) ([]byte, error)
}

type generator struct {
	packageName      string
	templateFilePath string
}

// thanks to https://www.socketloop.com/tutorials/golang-underscore-or-snake-case-to-camel-case-example
func snakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
	isToUpper := false

	for k, v := range inputUnderScoreStr {
		if k == 0 {
			camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return

}

var re = regexp.MustCompile(`\{(.*?)\}`)

func (g generator) Generate(tomlPathFile string) ([]byte, error) {
	tomlByteData, err := ioutil.ReadFile(tomlPathFile)
	if err != nil {
		return nil, err
	}

	rawStruct := map[string]interface{}{}

	if _, err = toml.Decode(string(tomlByteData), &rawStruct); err != nil {
		return nil, err
	}

	generator := GeneratedData{
		PackageName: g.packageName,
	}

	structChecker := make(map[string]bool)

	var structs []Struct
	for key := range rawStruct {
		structChecker[key] = true
	}

	for key, raw := range rawStruct {
		rawAsMap, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}

		structItem := Struct{
			StructName: key,
		}

		var fields []StructFields
		for rawKey, rawValue := range rawAsMap {
			t := reflect.TypeOf(rawValue)
			typeName := t.String()
			if structChecker[rawKey] {
				typeName = rawKey
				if t.String() == "[]map[string]interface {}" {
					typeName = "[]" + rawKey
				}

			} else {
				switch typeName {
				case "string":
					fieldValueString := rawValue.(string)
					_, err := time.Parse(time.RFC3339, fieldValueString)
					if err == nil {
						typeName = "time.Time"
						generator.UsingTime = true
					}

					valuesFromRegex := re.FindStringSubmatch(fieldValueString)
					if len(valuesFromRegex) > 1 {
						typeName = valuesFromRegex[1]
					}
				case "map[string]interface {}":
					typeName = "map[string]interface{}"
				case "[]interface {}":
					for _, value := range rawValue.([]interface{}) {
						t := reflect.TypeOf(value)
						typeName = "[]" + t.String()
						break
					}
				}
			}

			f := StructFields{
				Name:           snakeCaseToCamelCase(rawKey),
				Type:           typeName,
				JSONDescriptor: strings.ToLower(rawKey),
			}
			fields = append(fields, f)
		}
		structItem.Fields = fields
		structs = append(structs, structItem)
	}
	generator.Structs = structs

	tpl, err := template.ParseFiles(g.templateFilePath)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	if err = tpl.Execute(&buf, generator); err != nil {
		return nil, err
	}

	formattedString, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return formattedString, nil
}

func NewGenerator(packageName, templateFilePath string) Generator {
	return generator{
		packageName:      packageName,
		templateFilePath: templateFilePath,
	}
}
