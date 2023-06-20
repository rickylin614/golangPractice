package service

import (
	"nunuim/internal/dao"
	"nunuim/internal/model"
)

type GroupService struct {
	*Service
	groupDao *dao.GroupDao
}

func NewGroupService(service *Service, groupDao *dao.GroupDao) *GroupService {
	return &GroupService{
		Service: service,
		groupDao: groupDao,
	}
}

func (g *GroupService) GetGroupById(id int64) (*model.Group, error) {
	return g.groupDao.FirstById(id)
}