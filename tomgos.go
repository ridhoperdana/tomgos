package tomgos

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
)

type StructFields struct {
	Name string
	Type string
}

type GeneratedData struct {
	PackageName string
	StructName  string
	Fields      []StructFields
	UsingTime   bool
}

type Generator interface {
	Generate(tomlPathFile, targetFile string) error
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

func (g generator) Generate(tomlPathFile, targetFile string) error {
	tomlByteData, err := ioutil.ReadFile(tomlPathFile)
	if err != nil {
		return err
	}

	rawStruct := map[string]interface{}{}

	if _, err = toml.Decode(string(tomlByteData), &rawStruct); err != nil {
		return err
	}

	generator := GeneratedData{
		PackageName: g.packageName,
	}

	for key, raw := range rawStruct {
		generator.StructName = key
		rawAsMap, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		var fields []StructFields
		for rawKey, rawValue := range rawAsMap {
			t := reflect.TypeOf(rawValue)
			typeName := t.Name()
			if typeName == "string" {
				_, err := time.Parse(time.RFC3339, rawValue.(string))
				if err == nil {
					typeName = "time.Time"
					generator.UsingTime = true
				}
			}
			f := StructFields{
				Name: snakeCaseToCamelCase(rawKey),
				Type: typeName,
			}
			fields = append(fields, f)
		}
		generator.Fields = fields
	}

	tpl, err := template.ParseFiles(g.templateFilePath)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tpl.Execute(&buf, generator); err != nil {
		return err
	}

	formattedString, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(formattedString))
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		return err
	}

	return nil
}

func NewGenerator(packageName, templateFilePath string) Generator {
	return generator{
		packageName:      packageName,
		templateFilePath: templateFilePath,
	}
}
