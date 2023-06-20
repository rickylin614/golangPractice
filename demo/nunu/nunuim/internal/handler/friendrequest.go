package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/service"
	"nunuim/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type FriendRequestHandler struct {
	*Handler
	friendRequestService *service.FriendRequestService
}

func NewFriendRequestHandler(handler *Handler, friendRequestService *service.FriendRequestService) *FriendRequestHandler {
	return &FriendRequestHandler{
		Handler:     handler,
		friendRequestService: friendRequestService,
	}
}

func (f *FriendRequestHandler) GetFriendRequestById(ctx *gin.Context) {

	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	friendRequest, err := f.friendRequestService.GetFriendRequestById(params.Id)
	f.logger.Info("GetFriendRequestByID", zap.Any("friendRequest", friendRequest))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, friendRequest)
}
