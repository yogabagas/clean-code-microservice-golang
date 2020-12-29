package cmd

import (
	"my-github/clean-code-microservice-golang/config"
	"my-github/clean-code-microservice-golang/infrastructure/rest"

	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "api",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.WithField("component", "api_command").Infof("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		rest.NewRest(config.GetInt("rest.port"), db, rdb).Serve()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer db.Close()
	},
}
