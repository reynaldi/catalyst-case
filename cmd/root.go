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
	log.Println(os.Getenv("CONNECTION_STRING"))
	Conf = &config.Config{
		ConnectionString: os.Getenv("CONNECTION_STRING"),
		Dialect:          os.Getenv("DIALECT"),
		AppPort:          os.Getenv("APP_PORT"),
		AppHost:          os.Getenv("APP_HOST"),
	}
}
