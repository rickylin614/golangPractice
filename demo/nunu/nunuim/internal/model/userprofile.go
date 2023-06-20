package model

import "time"

type UserProfile struct {
	UserID      uint `gorm:"primaryKey"`
	AvatarURL   string
	DisplayName string
	Bio         string
	Location    string
	Website     string
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

// func (u *UserProfile) TableName() string {
// 	return "userProfile"
// }
