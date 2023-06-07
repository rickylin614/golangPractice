/* generate a golang project , use gin web framework, gorm, and have API: table user(account string, pwd string, age int, name string, nickname string) */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	Account  string `json:"account"`
	Pwd      string `json:"pwd"`
	Age      int    `json:"age"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}

func main() {
	r := gin.Default()
	r.GET("/user/:account", getUser)
	r.POST("/user", addUser)
	r.PUT("/user/:account", updateUser)
	r.DELETE("/user/:account", deleteUser)
	r.Run(":8080")
}

func getUser(c *gin.Context) {
	account := c.Param("account")
	var user User
	db.Where("account = ?", account).First(&user)
	c.JSON(http.StatusOK, user)
}

func addUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	account := c.Param("account")
	var user User
	db.Where("account = ?", account).First(&user)
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	account := c.Param("account")
	var user User
	db.Where("account = ?", account).First(&user)
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
