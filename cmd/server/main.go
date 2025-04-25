package main

import (
	"log"
	"net/http"
	"time"

	"github.com/CodeSync/internal/config"
)

func main() {
	cfg := config.Load()
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("server listening on : %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
