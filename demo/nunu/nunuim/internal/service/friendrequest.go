package service

import (
	"nunuim/internal/dao"
	"nunuim/internal/model"
)

type FriendRequestService struct {
	*Service
	friendRequestDao *dao.FriendRequestDao
}

func NewFriendRequestService(service *Service, friendRequestDao *dao.FriendRequestDao) *FriendRequestService {
	return &FriendRequestService{
		Service: service,
		friendRequestDao: friendRequestDao,
	}
}

func (f *FriendRequestService) GetFriendRequestById(id int64) (*model.FriendRequest, error) {
	return f.friendRequestDao.FirstById(id)
}