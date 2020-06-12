package cmd

import "github.com/spf13/cobra"

var (
	RootCMD = &cobra.Command{
		Use:   "tomgos",
		Short: "Tomgos is a helper to generate Go struct from TOML file",
		Long:  "For more information, see https://github.com/ridhoperdana/tomgos",
	}
)

func init() {
	RootCMD.AddCommand(generatorCMD)
	generatorCMD.Flags().StringVarP(&packageName, "package", "p", "tomgos",
		"Package name for the generated file. Default is tomgos")
	generatorCMD.Flags().StringVarP(&tomlFileLocation, "toml", "t", "file.toml",
		"TOML File location. Default is ./file.toml")
	generatorCMD.Flags().StringVarP(&targetFileLocation, "output", "o", "generated.go",
		"Target File location. Default is ./generated.go")
}
