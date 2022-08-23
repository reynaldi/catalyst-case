package server

import (
	"catalyst-case/config"
	"catalyst-case/database"
	"catalyst-case/internal/api"
	"fmt"
	"net/http"
)

type Server struct {
	*http.Server
	*config.Config
}

func NewServer(cfg *config.Config, db *database.DB) (*Server, error) {
	api := api.NewApi(cfg, db)
	server := api.CreateRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.AppHost, cfg.AppPort),
		Handler: server,
	}
	return &Server{srv, cfg}, nil
}
