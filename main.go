package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ithaquaKr/rssapp/pkg/config"
)

func main() {
	// Get .env file
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot get env variable, err: %s", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	addr := config.App.Host + ":" + config.App.Port
	fmt.Printf("Server is serving request at %s\n", addr)
	http.ListenAndServe(addr, r)

	// Setup database
	// db.NewDBConn()
}
