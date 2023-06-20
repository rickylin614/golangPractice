//go:build wireinject
// +build wireinject
package wire

import (
	"nunuim/internal/provider"
	"nunuim/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// wire.go 初始化模块
func NewApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		provider.ServerSet,
		provider.DaoSet,
		provider.ServiceSet,
		provider.HandlerSet,
		provider.SonyflakeSet,
		provider.JwtSet,
	))
}
