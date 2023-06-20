package service

import (
	"nunuim/internal/dao"
	"nunuim/internal/model"
)

type FriendService struct {
	*Service
	friendDao *dao.FriendDao
}

func NewFriendService(service *Service, friendDao *dao.FriendDao) *FriendService {
	return &FriendService{
		Service: service,
		friendDao: friendDao,
	}
}

func (f *FriendService) GetFriendById(id int64) (*model.Friend, error) {
	return f.friendDao.FirstById(id)
}