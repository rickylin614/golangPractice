package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/service"
	"nunuim/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type UserProfileHandler struct {
	*Handler
	userProfileService *service.UserProfileService
}

func NewUserProfileHandler(handler *Handler, userProfileService *service.UserProfileService) *UserProfileHandler {
	return &UserProfileHandler{
		Handler:     handler,
		userProfileService: userProfileService,
	}
}

func (u *UserProfileHandler) GetUserProfileById(ctx *gin.Context) {

	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	userProfile, err := u.userProfileService.GetUserProfileById(params.Id)
	u.logger.Info("GetUserProfileByID", zap.Any("userProfile", userProfile))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, userProfile)
}
