package cmd

import (
	"log"
	"os"

	"github.com/ridhoperdana/tomgos"
	"github.com/spf13/cobra"
)

var (
	packageName        = "tomgos"
	tomlFileLocation   = "file.toml"
	targetFileLocation = "generated.go"
	generatorCMD       = &cobra.Command{
		Use:   "generate",
		Short: "Read toml file and generate a Go file with struct written inside.",
		Run: func(cmd *cobra.Command, args []string) {
			generator := tomgos.NewGenerator(packageName, "template.txt")
			result, err := generator.Generate(tomlFileLocation)
			if err != nil {
				log.Fatal(err)
			}

			f, err := os.OpenFile(targetFileLocation, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatal(err)
			}

			_, err = f.WriteString(string(result))
			defer func() {
				if err := f.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			if err != nil {
				log.Fatal(err)
			}

			log.Println("Success generate the file")
		},
	}
)
