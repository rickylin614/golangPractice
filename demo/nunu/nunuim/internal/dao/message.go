package dao

import (
	"nunuim/internal/model"
)

type MessageDao struct {
	*Dao
}

func NewMessageDao(dao *Dao) *MessageDao {
	return &MessageDao{
		Dao: dao,
	}
}

func (m *MessageDao) FirstById(id int64) (*model.Message, error) {
	var message model.Message
	if err := m.db.Where("id = ?", id).First(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}