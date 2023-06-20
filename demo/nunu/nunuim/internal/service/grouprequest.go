package service

import (
	"nunuim/internal/dao"
	"nunuim/internal/model"
)

type GroupRequestService struct {
	*Service
	groupRequestDao *dao.GroupRequestDao
}

func NewGroupRequestService(service *Service, groupRequestDao *dao.GroupRequestDao) *GroupRequestService {
	return &GroupRequestService{
		Service: service,
		groupRequestDao: groupRequestDao,
	}
}

func (g *GroupRequestService) GetGroupRequestById(id int64) (*model.GroupRequest, error) {
	return g.groupRequestDao.FirstById(id)
}