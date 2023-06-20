package dao

import (
	"nunuim/internal/model"
)

type UserProfileDao struct {
	*Dao
}

func NewUserProfileDao(dao *Dao) *UserProfileDao {
	return &UserProfileDao{
		Dao: dao,
	}
}

func (u *UserProfileDao) FirstById(id int64) (*model.UserProfile, error) {
	var userProfile model.UserProfile
	if err := u.db.Where("id = ?", id).First(&userProfile).Error; err != nil {
		return nil, err
	}
	return &userProfile, nil
}