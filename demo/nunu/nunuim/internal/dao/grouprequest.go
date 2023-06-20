package dao

import (
	"nunuim/internal/model"
)

type GroupRequestDao struct {
	*Dao
}

func NewGroupRequestDao(dao *Dao) *GroupRequestDao {
	return &GroupRequestDao{
		Dao: dao,
	}
}

func (g *GroupRequestDao) FirstById(id int64) (*model.GroupRequest, error) {
	var groupRequest model.GroupRequest
	if err := g.db.Where("id = ?", id).First(&groupRequest).Error; err != nil {
		return nil, err
	}
	return &groupRequest, nil
}