package cmd

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "clean-code-microservice-golang",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof("root command")
	},
}

func Run() {
	rootCommand.AddCommand(restCommand)

	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
}
