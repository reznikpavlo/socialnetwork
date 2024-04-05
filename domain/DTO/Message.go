package DTO

import (
	"socialnetwork/domain/interfaces"
)

type Message struct {
	Id   int64  `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

func NewMessage(message interfaces.Message) Message {
	return Message{
		Id:   message.GetId(),
		Text: message.GetMessage(),
	}
}

func New() Message {
	return Message{}
}
func (m *Message) GetId() int64 {
	return m.Id
}

func (m *Message) GetMessage() string {
	return m.Text
}

func (m *Message) SetId(id int64) {
	m.Id = id
}

func (m *Message) SetText(t string) {
	m.Text = t
}
