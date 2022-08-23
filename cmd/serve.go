package cmd

import (
	"catalyst-case/database"
	"catalyst-case/pkg/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	db, err := database.Open(Conf)
	if err != nil {
		panic(err)
	}
	server, err := server.NewServer(Conf, db)
	if err != nil {
		panic(err)
	}

	go server.ListenAndServe()
	log.Printf("Listening on %v...", server.Addr)

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	sig := <-quit
	log.Println("Shutting down server... Reason: ", sig.String())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Print("server shutdown", err)
	}
	log.Println("Server gracefully stopped")
}
