package cmd

import (
	"catalyst-case/config"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "catalyst-base",
	Short: "Catalyst Test API",
	Long:  "Catalyst Test API",
}

var Conf *config.Config

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(loadConfig)
}

func loadConfig() {
	Conf.ConnectionString = os.Getenv("CONNECTION_STRING")
	Conf.Dialect = os.Getenv("DIALECT")
}
