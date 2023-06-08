package server

import (
	"net/http"
	"practice/demo/wire/service/bank"
	"practice/demo/wire/service/user"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type App struct {
	BankService bank.BankService
	UserService user.UserService
}

func NewApp() *App {
	wire.Build(wire.Struct(new(App), "*"),
		bank.NewBankService,
		user.NewUserService,
	)
	return &App{}
}

func (app App) Run() {
	r := gin.Default()

	r.GET("/bank", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, app.BankService.GetBank())
	})
	r.GET("/user", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, app.UserService.GetUser())
	})

	r.Run()
}
