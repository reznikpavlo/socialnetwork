package connectors

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
)

/*
social - name of database
Collections:
meeesages - contain messages
*/
type MongoDbConnector struct {
	source            string
	port              string //default 2701
	username          string
	password          string
	uri               string
	client            *mongo.Client
	MessageCollection *mongo.Collection
	Context           context.Context
}

func NewMongoDbConnector(source, port, username, password string) *MongoDbConnector {
	m := MongoDbConnector{
		source:            source,
		port:              port,
		username:          username,
		password:          password,
		uri:               "",
		MessageCollection: nil,
	}
	m.BuildUri()
	return &m
}
func (m *MongoDbConnector) BuildUri() {
	uriBuilder := strings.Builder{}
	uriBuilder.WriteString("mongodb://")
	//	uriBuilder.WriteString(m.username)
	//	uriBuilder.WriteString(":")
	//	uriBuilder.WriteString(m.password)
	//	uriBuilder.WriteString("@")
	uriBuilder.WriteString(m.source)
	uriBuilder.WriteString(":")
	uriBuilder.WriteString(m.port)
	//	uriBuilder.WriteString("/social")
	m.uri = uriBuilder.String()
	fmt.Println("uri do db = ", m.uri)
}

func (m *MongoDbConnector) ConnectDB() {
	if m.MessageCollection != nil {
		return
	}
	credential := options.Credential{Username: m.username, Password: m.password}
	clientOptions := options.Client()
	clientOptions.ApplyURI(m.uri).SetAuth(credential)
	client, err1 := mongo.Connect(context.Background(), clientOptions)
	err2 := client.Ping(m.Context, nil)
	log.Println("Connected to mongodb")
	if err2 != nil {
		log.Println(err2)
	}
	if err1 != nil {
		panic(err1)
	}
	m.MessageCollection = client.Database("social").Collection("messages")
	m.client = client

}

func (m *MongoDbConnector) Disconn() {
	m.client.Disconnect(m.Context)
}
