package model

type Accounts struct {
	Id          int    `gorm:"id" json:"id"`
	Name        string `gorm:"name" json:"name"`
	Pwd         string `gorm:"pwd" json:"pwd"`
	Testint     int    `gorm:"testint" json:"testint"`
	Testvarchar string `gorm:"testvarchar" json:"testvarchar"`
}

func (*Accounts) TableName() string {
	return "accounts"
}
