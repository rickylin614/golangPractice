package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	gouuid "github.com/gofrs/uuid"
	olduuid "github.com/google/uuid"
	"github.com/rickylin614/common/zlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

type Sf struct {
	Id   int64
	Data int64
}

type Uuid struct {
	Id   int
	Data string
}

func main() {

	count := 10
	sfnode, _ := snowflake.NewNode(1)

	db, err := InitDB("localhost", "demo", "root", "123456")
	if err != nil {
		panic(err)
	}

	var v4Du time.Duration
	var v6Du time.Duration
	var v7Du time.Duration
	var sfDu time.Duration

	for i := 0; i < 1; i++ {
		t7 := time.Now()
		v7arr := make([]Uuid, count)
		for i := range v7arr {
			v7, _ := gouuid.NewV7()
			v7arr[i].Id = 1
			v7arr[i].Data = strings.ReplaceAll(v7.String(), "-", "")
		}
		v7DB := db.Table("uuidv7").Create(v7arr)
		if v7DB.Error != nil {
			panic(v7DB.Error)
		}
		v7Du += time.Now().Sub(t7)

		// fmt.Println("v4 -----------------------------")
		t4 := time.Now()
		v4arr := make([]Uuid, count)
		for i := range v4arr {
			v4 := olduuid.NewString()
			v4arr[i].Id = 1
			v4arr[i].Data = strings.ReplaceAll(v4, "-", "")
		}
		// v4DB := db.Table("uuidv4").Create(v4arr)
		// if v4DB.Error != nil {
		// 	panic(v4DB.Error)
		// }
		v4Du += time.Now().Sub(t4)

		// fmt.Println("v6 -----------------------------")
		t6 := time.Now()
		v6arr := make([]Uuid, count)
		for i := range v6arr {
			v6, _ := gouuid.NewV6()
			v6arr[i].Id = 1
			v6arr[i].Data = strings.ReplaceAll(v6.String(), "-", "")
		}
		v6DB := db.Table("uuidv6").Create(v6arr)
		if v6DB.Error != nil {
			panic(v6DB.Error)
		}

		v6Du += time.Now().Sub(t6)

		// fmt.Println("snowflake -----------------------------")
		sft := time.Now()
		sfarr := make([]Sf, count)
		for i := range sfarr {
			sfid := sfnode.Generate()
			sfarr[i].Id = int64(math.Pow(2, 63) - 1)
			sfarr[i].Data = sfid.Int64()
		}
		sfdb := db.Table("snowflake").Create(sfarr)
		if sfdb.Error != nil {
			panic(sfdb.Error)
		}

		sfDu += time.Now().Sub(sft)

		// fmt.Println("v7 -----------------------------")

		// if i == 0 {
		fmt.Printf("v4Du:\t%d\t,v6Du:\t%d\t,v7Du:\t%d\t,sfDu:\t%d\t\n", v4Du.Milliseconds(), v6Du.Milliseconds(), v7Du.Milliseconds(), sfDu.Milliseconds())
		// }
	}
	fmt.Printf("v4Du:\t%d\t,v6Du:\t%d\t,v7Du:\t%d\t,sfDu:\t%d\t\n", v4Du.Milliseconds(), v6Du.Milliseconds(), v7Du.Milliseconds(), sfDu.Milliseconds())
}

func InitDB(host, schema, user, password string) (db *gorm.DB, err error) {
	dsn := user + ":" + password + "@tcp(" + host + ")/" + schema + "?charset=utf8mb4&loc=Local&parseTime=true&timeout=30s"

	// log初始化設定
	logger := zapgorm2.Logger{
		ZapLogger:                 zlog.GetLog(),
		LogLevel:                  logger.Error,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}

	// 連線
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	// 設置連接池數據
	sqlDB, err := gormdb.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)
	// 閒置連線的最大存在時間
	sqlDB.SetConnMaxIdleTime(time.Second * 25)
	// 連線的最大生存時間 確保連線可以被驅動安全關閉 官方建議小於五分鐘
	sqlDB.SetConnMaxLifetime(time.Second * 25)

	err = sqlDB.Ping()

	return gormdb, err
}
