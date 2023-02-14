package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Load()
	startServer(router.Build(), config.Port)
}

func startServer(router *mux.Router, port int) {
	fmt.Printf("Starting server at port [%d]\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
