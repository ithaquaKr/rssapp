package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ithaquaKr/rssapp/config"
)

const (
	maxHeaderBytes = 1 << 20
	ReadTimeout    = 10
	WriteTimeout   = 10
)

type Server struct {
	cfg *config.Config
	// TODO: Implement logger
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() {
	server := &http.Server{
		Addr:           s.cfg.App.Port,
		ReadTimeout:    time.Second * ReadTimeout,
		WriteTimeout:   time.Second * WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
		Handler:        test(),
	}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	go func(serverCtx context.Context) {
		fmt.Printf("Server is listening on :%s", s.cfg.App.Port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while starting Server: %s", err)
		}
		<-serverCtx.Done()
	}(serverCtx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	shutdownCtx, shutdown := context.WithTimeout(serverCtx, 5*time.Second)
	defer shutdown()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatalf("Error while shutdown server: %s", err)
	}
	serverStopCtx()
}

func test() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	return r
}
