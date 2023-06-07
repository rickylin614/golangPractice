package service

import (
	"demo/wire/service/bank"
	"demo/wire/service/user"

	"github.com/google/wire"
)

func InitService() {
	wire.Build(
		bank.NewBankService,
		user.NewUserService,
	)
}
