package main

import (
	"log"
	"net/http"
	"time"

	"github.com/CodeSync/internal/config"
	"github.com/CodeSync/internal/server"
)

func main() {
	cfg := config.Load()
	mux := server.NewRouter()

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("server listening on : %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
