package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"socialnetwork/domain"
	"strconv"
)

type Message struct {
	*http.ServeMux
}

var messages = []domain.Message{
	domain.NewMessage(1, "first message"),
	domain.NewMessage(2, "second message"),
	domain.NewMessage(3, "third message"),
	domain.NewMessage(4, "fourth message"),
}

func (c *Message) Get(w http.ResponseWriter, req *http.Request) {

	fmt.Println(messages)
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	getLogger := slog.New(jsonHandler)
	messagesJson, err := json.Marshal(messages)
	if err != nil {
		getLogger.Error("Message, GetMessages: error in getting all messages from db", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(messagesJson)
}

func (c *Message) Post(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Post method works")
	newMessage := domain.Message{}
	err := json.NewDecoder(req.Body).Decode(&newMessage)
	if err != nil {
		//TODO: log errors.
	}
	newMessage.Id = messages[len(messages)-1].Id + 1
	messages = append(messages, newMessage)
	messagesJson, _ := json.Marshal(messages)
	w.Write(messagesJson)
}

func (c *Message) Put(w http.ResponseWriter, req *http.Request) {
	value := req.PathValue("id")
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {

	}
	m := domain.Message{}
	errJson := json.NewDecoder(req.Body).Decode(&m)
	if errJson != nil {

	}
	fmt.Println(m)

	if id != m.Id {
		m.Id = id
	}

	messageFromDb := &messages[id-1]
	messageFromDb.Message = m.Message
	messagesJson, errJsonMarshal := json.Marshal(messages)
	if errJsonMarshal != nil {

	}
	w.Write(messagesJson)
}

func (c *Message) Delete(w http.ResponseWriter, req *http.Request) {
	value := req.PathValue("id")
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {

	}
	messages = append(messages[:id-1], messages[id:]...)
	messagesJson, _ := json.Marshal(messages)
	w.Write(messagesJson)
}

func (c *Message) RoutesInit() {
	c.ServeMux = http.NewServeMux()
	c.HandleFunc("GET /message", c.Get)
	c.HandleFunc("POST /message", c.Post)
	c.HandleFunc("PUT /message/{id}", c.Put)
	c.HandleFunc("DELETE /message/{id}", c.Delete)

}
