package service

import (
	"nunuim/internal/dao"
	"nunuim/internal/model"
)

type UserProfileService struct {
	*Service
	userProfileDao *dao.UserProfileDao
}

func NewUserProfileService(service *Service, userProfileDao *dao.UserProfileDao) *UserProfileService {
	return &UserProfileService{
		Service: service,
		userProfileDao: userProfileDao,
	}
}

func (u *UserProfileService) GetUserProfileById(id int64) (*model.UserProfile, error) {
	return u.userProfileDao.FirstById(id)
}