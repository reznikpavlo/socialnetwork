package connectors

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
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
	UserCollection    *mongo.Collection
	Context           context.Context
}

func NewMongoDbConnector() *MongoDbConnector {
	err := godotenv.Load("application.env")
	if err != nil {
		log.Println("error oppening env file.")
	}
	connector := newMongoDbConnector(os.Getenv("mongo.datasource.url"),
		os.Getenv("mongo.datasource.port"),
		os.Getenv("mongo.datasource.username"),
		os.Getenv("mongo.datasource.password"))
	return connector
	/*	connector := newMongoDbConnector("localhost",
			"27017",
			"admin",
			"Qwertyu8")
		return connector
	*/
}

func newMongoDbConnector(source, port, username, password string) *MongoDbConnector {
	m := MongoDbConnector{
		source:            source,
		port:              port,
		username:          username,
		password:          password,
		uri:               "",
		MessageCollection: nil,
		UserCollection:    nil,
		Context:           context.TODO(),
	}
	m.BuildUri()
	return &m
}
func (m *MongoDbConnector) BuildUri() {
	uriBuilder := strings.Builder{}
	uriBuilder.WriteString("mongodb://")
	uriBuilder.WriteString(m.source)
	uriBuilder.WriteString(":")
	uriBuilder.WriteString(m.port)
	m.uri = uriBuilder.String()
	//	log.Println("uri do db = ", m.uri)
}

func (m *MongoDbConnector) ConnectDB() {
	if m.MessageCollection != nil {
		return
	}
	credential := options.Credential{Username: m.username, Password: m.password}
	clientOptions := options.Client()
	clientOptions.ApplyURI(m.uri).SetAuth(credential)
	client, err1 := mongo.Connect(context.TODO(), clientOptions)
	err2 := client.Ping(m.Context, nil)
	log.Println("Connected to mongodb")
	if err2 != nil {
		log.Println(err2)
	}
	if err1 != nil {
		panic(err1)
	}
	m.MessageCollection = client.Database("social").Collection("messages")
	m.UserCollection = client.Database("social").Collection("usr")
	m.client = client

}

func (m *MongoDbConnector) Disconn() {
	m.client.Disconnect(m.Context)
}
