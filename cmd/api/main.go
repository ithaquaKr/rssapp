package main

import (
	"log"

	"github.com/ithaquaKr/rssapp/internal/server"
	"github.com/ithaquaKr/rssapp/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot get env variable, err: %s", err)
	}
	server := server.NewServer(config)
	server.Run()
}
