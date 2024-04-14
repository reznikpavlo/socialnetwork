package domain

import "time"

type Message struct {
	Id           int64     `json:"id"`
	Text         string    `json:"text"`
	CreationDate time.Time `json:"creationDate"`
}

func NewMessage(idd int64, mm string) *Message {
	return &Message{Id: idd, Text: mm, CreationDate: time.Now()}
}

func (m *Message) getDate() string {
	return m.CreationDate.Format("02-01-2006 15:04")
}
