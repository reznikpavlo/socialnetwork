package main

import (
	"log"
	"net/http"
	"socialnetwork/controller"
)

func main() {
	messageController := &controller.Message{}
	messageController.RoutesInit()

	mux := http.NewServeMux()
	mux.Handle("/message/", messageController.Serve.ServeMux)
	mux.Handle("/", http.FileServer(http.Dir("./templates/static")))

	log.Println("Server running on: localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: mux}
	server.ListenAndServe()

}
