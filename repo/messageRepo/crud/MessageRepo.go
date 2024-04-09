package crud

import "socialnetwork/domain"

type MessageRepo interface {
	FindById(id int64) domain.Message
	FindAll() []domain.Message
	FindTop(top int64) []domain.Message
	Save(m *domain.Message) domain.Message
	DeleteOne(m *domain.Message)
}
