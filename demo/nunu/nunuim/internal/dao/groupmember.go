package dao

import (
	"nunuim/internal/model"
)

type GroupMemberDao struct {
	*Dao
}

func NewGroupMemberDao(dao *Dao) *GroupMemberDao {
	return &GroupMemberDao{
		Dao: dao,
	}
}

func (g *GroupMemberDao) FirstById(id int64) (*model.GroupMember, error) {
	var groupMember model.GroupMember
	if err := g.db.Where("id = ?", id).First(&groupMember).Error; err != nil {
		return nil, err
	}
	return &groupMember, nil
}