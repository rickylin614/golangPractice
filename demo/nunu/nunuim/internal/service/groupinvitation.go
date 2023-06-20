package service

import (
	"nunuim/internal/dao"
	"nunuim/internal/model"
)

type GroupInvitationService struct {
	*Service
	groupInvitationDao *dao.GroupInvitationDao
}

func NewGroupInvitationService(service *Service, groupInvitationDao *dao.GroupInvitationDao) *GroupInvitationService {
	return &GroupInvitationService{
		Service: service,
		groupInvitationDao: groupInvitationDao,
	}
}

func (g *GroupInvitationService) GetGroupInvitationById(id int64) (*model.GroupInvitation, error) {
	return g.groupInvitationDao.FirstById(id)
}