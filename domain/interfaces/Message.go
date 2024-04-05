package interfaces

type Message interface {
	GetId() int64
	GetMessage() string
	SetId(id int64)
	SetText(t string)
}
