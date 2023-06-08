//go:build wireinject
// +build wireinject

package wire

import (
	"nunuim/internal/migration"
	"nunuim/internal/provider"
	"nunuim/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// wire.go 初始化模块
func NewApp(*viper.Viper, *log.Logger) (*migration.Migrate, func(), error) {
	//log.Info("NewApp")
	panic(wire.Build(
		provider.DaoSet,
		provider.MigrateSet,
	))
}
