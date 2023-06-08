package provider

import (
	"nunuim/internal/dao"
	"nunuim/internal/handler"
	"nunuim/internal/job"
	"nunuim/internal/middleware"
	"nunuim/internal/migration"
	"nunuim/internal/server"
	"nunuim/internal/service"
	"nunuim/pkg/helper/sonyflake"
	"github.com/google/wire"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var SonyflakeSet = wire.NewSet(sonyflake.NewSonyflake)

var MigrateSet = wire.NewSet(migration.NewMigrate)

var JobSet = wire.NewSet(job.NewJob)

var JwtSet = wire.NewSet(middleware.NewJwt)

var DaoSet = wire.NewSet(
	dao.NewDB,
	dao.NewRedis,
	dao.NewDao,
	dao.NewUserDao,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)
