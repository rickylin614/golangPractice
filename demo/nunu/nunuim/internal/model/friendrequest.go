package model

import "gorm.io/gorm"

type FriendRequest struct {
	gorm.Model
}

// func (f *FriendRequest) TableName() string {
// 	return "friendRequest"
// }
