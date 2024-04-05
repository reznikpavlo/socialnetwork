package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialnetwork/domain"
	"socialnetwork/domain/DTO"
	"strconv"
)

type Message struct {
	Serve Controller
}

var messages = []DTO.Message{
	DTO.NewMessage(domain.NewMessage(1, "first message")),
	DTO.NewMessage(domain.NewMessage(2, "second message")),
	DTO.NewMessage(domain.NewMessage(3, "third message")),
	DTO.NewMessage(domain.NewMessage(4, "fourth message")),
}

func (c *Message) Get(w http.ResponseWriter, req *http.Request) {

	messagesJson, err := json.Marshal(messages)
	if err != nil {

		//getLogger.Error("Text, GetMessages: error in getting all messages from db", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(messagesJson)
}

func (c *Message) Post(w http.ResponseWriter, req *http.Request) {

	newMessage := DTO.Message{}

	err := json.NewDecoder(req.Body).Decode(&newMessage)

	if err != nil {
		//TODO: log errors.
	}
	newMessage.SetId(messages[len(messages)-1].Id + 1)

	message := DTO.NewMessage(&newMessage)
	messages = append(messages, message)
	messagesJson, _ := json.Marshal(message)
	w.Write(messagesJson)
}

func (c *Message) Put(w http.ResponseWriter, req *http.Request) {
	value := req.PathValue("id")
	fmt.Println(value)
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {

	}
	m := DTO.Message{}
	errJson := json.NewDecoder(req.Body).Decode(&m)
	if errJson != nil {

	}
	fmt.Println(m)

	if id != m.GetId() {
		m.SetId(id)
	}

	messageFromDb := &messages[id-1]
	messageFromDb.Text = m.GetMessage()
	messagesJson, errJsonMarshal := json.Marshal(messageFromDb)
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

func (c *Message) GetOne(w http.ResponseWriter, req *http.Request) {
	value := req.PathValue("id")
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {

	}

	messageJson, errJson := json.Marshal(messages[id-1])
	if errJson != nil {
	}
	w.Write(messageJson)
}

func (c *Message) RoutesInit() {
	c.Serve.ServeMux = http.NewServeMux()
	c.Serve.HandleFunc("GET /message/", c.Get)
	c.Serve.HandleFunc("GET /message/{id}", c.GetOne)
	c.Serve.HandleFunc("POST /message/", c.Post)
	c.Serve.HandleFunc("PUT /message/{id}", c.Put)
	c.Serve.HandleFunc("DELETE /message/{id}", c.Delete)
}
