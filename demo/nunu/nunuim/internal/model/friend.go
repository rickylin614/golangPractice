package model

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
}

// func (f *Friend) TableName() string {
// 	return "friend"
// }
