package main

import (
	"log"
	"net/http"

	"github.com/nchukkaio/simchat/internal/handlers"
)

func main() {
	routes := routes()
	log.Println("starting channel listener")
	go handlers.ListenToWsChannel()
	log.Println("starting web server on port 8080")

	_ = http.ListenAndServe(":8080", routes)
}
