package main

import (
	"log"
	"net/http"
	"socialnetwork/controller"
	"socialnetwork/repo/dbconnections"
	"socialnetwork/repo/dbconnections/connectors"
	"socialnetwork/repo/messageRepo"
	"socialnetwork/service"
)

func main() {
	messageController := &controller.Message{}

	connector := connectors.NewMongoDbConnector("localhost", "27017", "admin", "Qwertyu8")
	mrepoMongo := messageRepo.MessageRepoMongo{
		&dbconnections.Db{MongoDB: connector},
	}
	messageController.Service = &service.Message{
		Repo: &mrepoMongo,
	}
	messageController.RoutesInit()

	mux := http.NewServeMux()
	mux.Handle("/message/", messageController.Serve.ServeMux)
	mux.Handle("/", http.FileServer(http.Dir("./templates/static")))

	log.Println("Server running on: localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: mux}
	server.ListenAndServe()

	defer server.Close()
	defer connector.Disconn()

}
