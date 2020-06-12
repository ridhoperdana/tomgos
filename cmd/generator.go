package cmd

import (
	"log"

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
			if err := generator.Generate(tomlFileLocation, targetFileLocation); err != nil {
				log.Fatal(err)
			}
		},
	}
)
