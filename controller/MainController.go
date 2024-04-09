package controller

import (
	"encoding/json"
	"net/http"
	"socialnetwork/domain"
	"socialnetwork/service"
	"strconv"
)

type Message struct {
	Serve   Controller
	Service *service.Message
}

func (c *Message) Get(w http.ResponseWriter, req *http.Request) {
	messages := c.Service.GetAllMessages()

	messagesJson, err := json.Marshal(messages)
	if err != nil {

		//getLogger.Error("Text, GetMessages: error in getting all messages from db", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(messagesJson)
}

func (c *Message) Post(w http.ResponseWriter, req *http.Request) {

	newMessage := domain.Message{}

	err := json.NewDecoder(req.Body).Decode(&newMessage)

	if err != nil {
		//TODO: log errors.
	}
	newMessage.Id = 0
	c.Service.Save(&newMessage)
	messagesJson, _ := json.Marshal(newMessage)
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

	if id != m.Id {
		m.Id = id
	}

	messageFromDb := c.Service.Save(&m)
	messagesJson, errJsonMarshal := json.Marshal(messageFromDb)
	if errJsonMarshal != nil {

	}
	w.Write(messagesJson)
}

func (c *Message) Delete(w http.ResponseWriter, req *http.Request) {
	value := req.PathValue("id")
	c.Service.DeleteOneById(value)
	messages := c.Service.GetAllMessages()
	messagesJson, _ := json.Marshal(messages)
	w.Write(messagesJson)
}

func (c *Message) GetOne(w http.ResponseWriter, req *http.Request) {
	value := req.PathValue("id")
	messageFromDb := c.Service.FindOneById(value)
	messageJson, errJson := json.Marshal(messageFromDb)
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
