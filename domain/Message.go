package domain

type Message struct {
	Id      int64  `json:"id,omitempty"`
	Message string `json:"message"`
}

func NewMessage(idd int64, mm string) Message {
	return Message{Id: idd, Message: mm}

}
