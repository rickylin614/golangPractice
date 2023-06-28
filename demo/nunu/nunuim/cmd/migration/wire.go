//go:build wireinject
// +build wireinject

package main

import (
	"nunuim/internal/migration"
	"nunuim/internal/repository"
	"nunuim/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)
var MigrateSet = wire.NewSet(migration.NewMigrate)

func newApp(*viper.Viper, *log.Logger) (*migration.Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		MigrateSet,
	))
}
