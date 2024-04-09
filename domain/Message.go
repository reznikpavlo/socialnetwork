package domain

type Message struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

func NewMessage(idd int64, mm string) *Message {
	return &Message{Id: idd, Text: mm}
}
