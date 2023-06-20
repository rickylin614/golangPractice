package dao

import (
	"nunuim/internal/model"
)

type GroupInvitationDao struct {
	*Dao
}

func NewGroupInvitationDao(dao *Dao) *GroupInvitationDao {
	return &GroupInvitationDao{
		Dao: dao,
	}
}

func (g *GroupInvitationDao) FirstById(id int64) (*model.GroupInvitation, error) {
	var groupInvitation model.GroupInvitation
	if err := g.db.Where("id = ?", id).First(&groupInvitation).Error; err != nil {
		return nil, err
	}
	return &groupInvitation, nil
}