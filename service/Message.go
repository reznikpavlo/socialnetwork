package service

import (
	"socialnetwork/domain"
	"socialnetwork/repo/messageRepo/crud"
	"strconv"
)

type Message struct {
	Repo crud.MessageRepo
}

func (service *Message) GetAllMessages() []domain.Message {
	messages := service.Repo.FindAll()
	return messages
}

func (service *Message) GetTop(top int64) []domain.Message {
	messages := service.Repo.FindTop(top)
	return messages
}

func (service *Message) Save(m *domain.Message) domain.Message {
	return service.Repo.Save(m)
}

func (service *Message) DeleteOne(message *domain.Message) {
	service.Repo.DeleteOne(message)
}

func (service *Message) DeleteOneById(id string) {
	idd, _ := strconv.ParseInt(id, 10, 64)
	message := service.Repo.FindById(idd)
	service.Repo.DeleteOne(&message)
}

func (m *Message) FindOneById(id string) *domain.Message {
	idd, _ := strconv.ParseInt(id, 10, 64)
	message := m.Repo.FindById(idd)
	return &message
}
