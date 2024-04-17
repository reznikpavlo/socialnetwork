package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"socialnetwork/controller"
	"socialnetwork/repo/SessionRepo/implementation"
	"socialnetwork/repo/dbconnections"
	"socialnetwork/repo/dbconnections/connectors"
	"socialnetwork/repo/messageRepo"
	"socialnetwork/repo/userRepo"
	"socialnetwork/service"
	"socialnetwork/utils"
	"syscall"
	"time"
)

func main() {
	//initializing mongodb
	connector := connectors.NewMongoDbConnector()
	mrepoMongo := messageRepo.MessageRepoMongo{
		&dbconnections.Db{MongoDB: connector},
	}

	//message controller init
	messageController := &controller.Message{}
	messageController.Service = &service.Message{
		Repo: &mrepoMongo,
	}
	messageController.RoutesInit()
	// initializing users repo and service
	urepo := &userRepo.UserRepo{&dbconnections.Db{MongoDB: connector}}
	sessionRepo := implementation.RepoInit()
	// sesionserv := service.NewSessionService(sessionRepo, 86000)
	userService := service.NewUserService(urepo, sessionRepo)
	// login initializing
	loginController := &controller.LoginController{}
	loginController.Init(userService)

	mux := http.NewServeMux()
	mux.Handle("/message/", messageController.Serve.ServeMux)
	mux.Handle("/google/", loginController.Serve.ServeMux)
	mux.Handle("/", http.FileServer(http.Dir("./templates/static/")))

	uri := utils.NewServerAddress()
	server := &http.Server{Addr: uri, Handler: mux}
	log.Println("Server running on ", uri)
	err := server.ListenAndServe()

	defer server.Close()
	defer connector.Disconn()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server ListenAndServe error: %v", err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown error: %v", err)
	}
	log.Println("Server gracefully stopped.")

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
