package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server",
	Long:  "start http server withw configured API",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func run() {
	log.Println("Creating new server")
}
