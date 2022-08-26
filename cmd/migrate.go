package cmd

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	m, err := openConnection()
	if err != nil {
		log.Println("open connection error")
		log.Println(err)
		return err
	}
	err = m.Up()
	if err != nil {
		log.Println("migration error")
		log.Println(err)
		return err
	}
	log.Println(m.Version())

	return nil
}

func openConnection() (*migrate.Migrate, error) {
	db, err := sql.Open(Conf.Dialect, Conf.ConnectionString)
	if err != nil {
		log.Printf("error when opening db connection: %v", err)
		return nil, err
	}
	defer db.Close()
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"mysql", driver)

	if err != nil {
		log.Println("error create migration")
		return nil, err
	}
	return m, nil
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
