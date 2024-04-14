package main

import (
	"fmt"
	"log"
	"net/http"
	"socialnetwork/controller"
	"socialnetwork/repo/dbconnections"
	"socialnetwork/repo/dbconnections/connectors"
	"socialnetwork/repo/messageRepo"
	"socialnetwork/service"
	"socialnetwork/utils"
)

func main() {
	messageController := &controller.Message{}
	connector := connectors.NewMongoDbConnector()
	mrepoMongo := messageRepo.MessageRepoMongo{
		&dbconnections.Db{MongoDB: connector},
	}
	messageController.Service = &service.Message{
		Repo: &mrepoMongo,
	}
	messageController.RoutesInit()
	loginController := &controller.LoginController{}
	loginController.Init()
	mux := http.NewServeMux()
	mux.Handle("/message/", messageController.Serve.ServeMux)
	mux.Handle("/google/", loginController.Serve.ServeMux)
	mux.Handle("/", http.FileServer(http.Dir("./templates/static")))
	uri := utils.NewServerAddress()
	server := &http.Server{Addr: uri, Handler: mux}
	log.Println("Server running on ", uri)

	err := server.ListenAndServe()
	log.Println(err)
	checkErr(err)
	defer server.Close()
	defer connector.Disconn()

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
