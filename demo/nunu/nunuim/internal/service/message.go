package service

import (
	"nunuim/internal/dao"
	"nunuim/internal/model"
)

type MessageService struct {
	*Service
	messageDao *dao.MessageDao
}

func NewMessageService(service *Service, messageDao *dao.MessageDao) *MessageService {
	return &MessageService{
		Service: service,
		messageDao: messageDao,
	}
}

func (m *MessageService) GetMessageById(id int64) (*model.Message, error) {
	return m.messageDao.FirstById(id)
}