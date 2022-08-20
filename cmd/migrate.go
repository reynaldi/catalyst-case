package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate db",
	Long:  "Migrate db",
	Run: func(cmd *cobra.Command, args []string) {
		var err = migrateDb()
		if err != nil {
			log.Printf("error on migration : %v", err)
		}
	},
}

func migrateDb() error {
	return nil
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
