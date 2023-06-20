package dao

import (
	"nunuim/internal/model"
)

type GroupDao struct {
	*Dao
}

func NewGroupDao(dao *Dao) *GroupDao {
	return &GroupDao{
		Dao: dao,
	}
}

func (g *GroupDao) FirstById(id int64) (*model.Group, error) {
	var group model.Group
	if err := g.db.Where("id = ?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}