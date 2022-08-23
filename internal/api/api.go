package api

import (
	"catalyst-case/config"
	"catalyst-case/database"
	"fmt"
	"net/http"
)

type Api struct {
	cfg *config.Config
	db  *database.DB
}

type Router interface {
	CreateRouter() *http.ServeMux
}

func NewApi(cfg *config.Config, db *database.DB) *Api {
	return &Api{
		cfg: cfg,
		db:  db,
	}
}

func (api *Api) CreateRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
