package service

import (
	"practice/demo/wire/service/bank"
	"practice/demo/wire/service/user"

	"github.com/google/wire"
)

func InitService() {
	wire.Build(
		bank.NewBankService,
		user.NewUserService,
	)
}
