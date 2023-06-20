package dao

import (
	"nunuim/internal/model"
)

type FriendDao struct {
	*Dao
}

func NewFriendDao(dao *Dao) *FriendDao {
	return &FriendDao{
		Dao: dao,
	}
}

func (f *FriendDao) FirstById(id int64) (*model.Friend, error) {
	var friend model.Friend
	if err := f.db.Where("id = ?", id).First(&friend).Error; err != nil {
		return nil, err
	}
	return &friend, nil
}