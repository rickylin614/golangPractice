package main

import (
	"fmt"

	"github.com/rickylin614/common/cgorm"
	"gorm.io/gorm"
)

// User has many CreditCards, UserID is the foreign key
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint
}

type AAAA struct {
	gorm.Model
	Name   string `gorm:"name"` // ,omitempty
	Age    int    `gorm:"age"`
	BsName string `gorm:"bs_name"`
	Bs     []BBBB `gorm:"foreignKey:Name;references:bs_name"`
}

type BBBB struct {
	ID   string `gorm:"primarykey"`
	Name string
}

func main() {
	cgorm.InitDB("localhost", "demo", "root", "123456", "")

	db := cgorm.GetDB()
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&CreditCard{})
	// db.AutoMigrate(&AAAA{})
	// db.AutoMigrate(&BBBB{})
	// db = cgorm.Begin()
	// db = db.Create(&AAAA{Model: gorm.Model{ID: 5}, Name: "string", Age: 20})

	var users []User
	db.Model(&User{}).Preload("CreditCards").Find(&users)
	fmt.Println(users)

	str := ToSQL(db, func(tx *gorm.DB) *gorm.DB {
		var users []User
		tx = tx.Joins("left join CreditCard cc on cc.UserRefer = User.Id").Find(&users)
		return tx
	})

	fmt.Println(str)
	// err := db.Commit()

	fmt.Println("end")

}

func ToSQL(db *gorm.DB, queryFn func(tx *gorm.DB) *gorm.DB) string {
	tx := queryFn(db.Session(&gorm.Session{DryRun: true, SkipDefaultTransaction: true}))
	stmt := tx.Statement
	return db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
}
