package main

import (
	"fmt"
	"github.com/gayibovih/http-server/config"
	"github.com/gayibovih/http-server/internal/server"
	"log"
	"net/http"
)

func main() {
	cfg := config.GetConfig()

	srv := server.New()

	log.Printf("Server run on %d port", cfg.HttpPort)
	ListenAddr := fmt.Sprintf(":%d", cfg.HttpPort)
	http.ListenAndServe(ListenAddr, srv.Router)

}
