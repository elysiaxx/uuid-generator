package cmd

import (
	"fmt"
	"github.com/elysiaxx/uuid-generator/generator"

	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Check version of uuid",
	Long: `Check version of uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GetVersion())
	},
}

func GetVersion() string {
	g := generator.GeneratorV4{}
	return g.Version()
}