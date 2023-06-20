package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/service"
	"nunuim/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type FriendHandler struct {
	*Handler
	friendService *service.FriendService
}

func NewFriendHandler(handler *Handler, friendService *service.FriendService) *FriendHandler {
	return &FriendHandler{
		Handler:     handler,
		friendService: friendService,
	}
}

func (f *FriendHandler) GetFriendById(ctx *gin.Context) {

	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	friend, err := f.friendService.GetFriendById(params.Id)
	f.logger.Info("GetFriendByID", zap.Any("friend", friend))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, friend)
}
