package service

import (
	"wiretest/service/bank"
	"wiretest/service/user"

	"github.com/google/wire"
)

func InitService() {
	wire.Build(
		bank.NewBankService,
		user.NewUserService,
	)
}
