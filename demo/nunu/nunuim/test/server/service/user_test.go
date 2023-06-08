package service

import (
	"fmt"
	"nunuim/internal/dao"
	"nunuim/internal/middleware"
	"nunuim/internal/service"
	"nunuim/pkg/config"
	"nunuim/pkg/helper/sonyflake"
	"nunuim/pkg/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var userService *service.UserService

func TestMain(m *testing.M) {
	fmt.Println("begin")

	os.Setenv("APP_CONF", "../../../config/local.yml")

	conf := config.NewConfig()

	logger := log.NewLog(conf)
	db := dao.NewDB(conf)
	rdb := dao.NewRedis(conf)
	jwt := middleware.NewJwt(conf)
	sf := sonyflake.NewSonyflake()
	srv := service.NewService(logger, sf, jwt)
	repo := dao.NewDao(db, rdb, logger)
	userDao := dao.NewUserDao(repo)
	userService = service.NewUserService(srv, userDao)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}
func TestRegister(t *testing.T) {
	req := service.RegisterRequest{
		Username: "user1",
		Password: "123456",
		Email:    "user1@mail.com",
	}
	err := userService.Register(&req)
	assert.Equal(t, err, nil, "they should be equal")
}

func TestLogin(t *testing.T) {
	req := service.LoginRequest{
		Username: "user1",
		Password: "123456",
	}
	token, err := userService.Login(&req)
	assert.Equal(t, err, nil, "they should be equal")
	t.Log("token", token)
}
