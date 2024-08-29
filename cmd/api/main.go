package main

import (
	"log"

	"github.com/ithaquaKr/rssapp/config"
	"github.com/ithaquaKr/rssapp/internal/server"
)

func main() {
	log.Println("Starting api server")
	config, err := config.InitConfig(".", "local-conf")
	if err != nil {
		log.Fatalf("Cannot get env variable, err: %s", err)
	}
	server := server.NewServer(config)
	server.Run()
}
