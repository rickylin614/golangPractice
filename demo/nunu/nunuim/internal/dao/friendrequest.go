package dao

import (
	"nunuim/internal/model"
)

type FriendRequestDao struct {
	*Dao
}

func NewFriendRequestDao(dao *Dao) *FriendRequestDao {
	return &FriendRequestDao{
		Dao: dao,
	}
}

func (f *FriendRequestDao) FirstById(id int64) (*model.FriendRequest, error) {
	var friendRequest model.FriendRequest
	if err := f.db.Where("id = ?", id).First(&friendRequest).Error; err != nil {
		return nil, err
	}
	return &friendRequest, nil
}