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
	dao.NewFriendDao,          // 添加 FriendDao
	dao.NewFriendRequestDao,   // 添加 FriendRequestDao
	dao.NewGroupDao,           // 添加 GroupDao
	dao.NewGroupInvitationDao, // 添加 GroupInvitationDao
	dao.NewGroupMemberDao,     // 添加 GroupMemberDao
	dao.NewGroupRequestDao,    // 添加 GroupRequestDao
	dao.NewMessageDao,         // 添加 MessageDao
	dao.NewUserProfileDao,     // 添加 UserProfileDao
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewFriendService,          // 添加 FriendService
	service.NewFriendRequestService,   // 添加 FriendRequestService
	service.NewGroupService,           // 添加 GroupService
	service.NewGroupInvitationService, // 添加 GroupInvitationService
	service.NewGroupMemberService,     // 添加 GroupMemberService
	service.NewGroupRequestService,    // 添加 GroupRequestService
	service.NewMessageService,         // 添加 MessageService
	service.NewUserProfileService,     // 添加 UserProfileService
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewFriendHandler,          // 添加 FriendHandler
	handler.NewFriendRequestHandler,   // 添加 FriendRequestHandler
	handler.NewGroupHandler,           // 添加 GroupHandler
	handler.NewGroupInvitationHandler, // 添加 GroupInvitationHandler
	handler.NewGroupMemberHandler,     // 添加 GroupMemberHandler
	handler.NewGroupRequestHandler,    // 添加 GroupRequestHandler
	handler.NewMessageHandler,         // 添加 MessageHandler
	handler.NewUserProfileHandler,     // 添加 UserProfileHandler
)
