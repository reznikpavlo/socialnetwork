package domain

import "socialnetwork/domain/interfaces"

type Message struct {
	id   int64
	text string
}

func NewMessage(idd int64, mm string) interfaces.Message {
	return &Message{id: idd, text: mm}
}

func (m *Message) GetId() int64 {
	return m.id
}

func (m *Message) GetMessage() string {
	return m.text
}

func (m *Message) SetId(id int64) {
	m.id = id
}

func (m *Message) SetText(text string) {
	m.text = text
}
