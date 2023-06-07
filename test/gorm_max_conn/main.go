package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DemoData struct {
	gorm.Model

	Name   string
	Age    int
	Remark string
}

var gormDB *gorm.DB

func (DemoData) TableName() string {
	return "demo_data"
}

func init() {
	err := InitDB("localhost", "demo", "root", "123456", "")
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&DemoData{})
}

func InitDB(host, schema, user, password, dbSourceName string) (err error) {
	dsn := user + ":" + password + "@tcp(" + host + ")/" + schema + "?charset=utf8mb4&loc=Local&parseTime=true&timeout=30s"
	gormDB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(0)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(1)
	// 閒置連線的最大存在時間
	sqlDB.SetConnMaxIdleTime(time.Second * 5)
	// 連線的最大生存時間 確保連線可以被驅動安全關閉 官方建議小於五分鐘
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	err = sqlDB.Ping()
	return
}

func main() {
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})

	connCount := 0

	r.GET("/demo/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		db := gormDB
		db = db.Begin()
		var demoData *DemoData
		db = db.Where("name", name).Find(demoData)
		if demoData == nil {
			demoData = &DemoData{
				Name:   name,
				Age:    rand.Intn(100),
				Remark: uuid.NewString(),
			}
			db = db.Create(demoData)
		}
		time.Sleep(30 * time.Second)
		connCount++
		db.Commit()
		n := time.Now().Format(time.StampMilli)
		fmt.Printf("%d: [%s] %s commmit\n", connCount, n, name)
		ctx.JSON(http.StatusOK, demoData)
	})

	r.Run()
}
