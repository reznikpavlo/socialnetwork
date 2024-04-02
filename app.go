package main

import (
	"net/http"
	"socialnetwork/controller"
)

func main() {
	messageController := &controller.Message{}
	messageController.RoutesInit()
	server := &http.Server{Addr: ":8080", Handler: messageController.ServeMux}
	server.ListenAndServe()
}
