package service

import (
	"nunuim/internal/dao"
	"nunuim/internal/model"
)

type GroupMemberService struct {
	*Service
	groupMemberDao *dao.GroupMemberDao
}

func NewGroupMemberService(service *Service, groupMemberDao *dao.GroupMemberDao) *GroupMemberService {
	return &GroupMemberService{
		Service: service,
		groupMemberDao: groupMemberDao,
	}
}

func (g *GroupMemberService) GetGroupMemberById(id int64) (*model.GroupMember, error) {
	return g.groupMemberDao.FirstById(id)
}