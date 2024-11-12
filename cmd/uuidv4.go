package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"github.com/elysiaxx/uuid-generator/generator"

	"github.com/spf13/cobra"
)

var uuidv4Cmd = &cobra.Command{
	Use: "uuidv4",
	Short: "Generate uuid version 4",
	Long: `Generate an uuid version 4 or the list of uuidv4 in number range 1 - 100`,
	Run: func(cmd *cobra.Command, args []string) {
		// parse args to get number of uuidv4 that need to produce
		gv4 := generator.GeneratorV4{}
		num, err := ParseArgs(args)
		if err != nil {
			fmt.Println("Fail to parse Args: ", err)
			os.Exit(1)
		}
		if num >= 100 {
			fmt.Println("the tool currently support to generate maximum of 100 uuids")
			os.Exit(1)
		}
		for range num {
			res, err := gv4.UUID()
			if err != nil {
				fmt.Printf("Generating UUIDs failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(gv4.String(res))
		}
	},
}

func ParseArgs(args []string) (uint, error) {
	if len(args) == 0 {
		return 1, nil
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errors.New("argument must be an integer")
	}
	if num < 1 {
		return 0, errors.New("please enter a number greater than 0")
	}
	return uint(num), nil
}