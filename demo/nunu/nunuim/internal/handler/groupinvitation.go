package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/service"
	"nunuim/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type GroupInvitationHandler struct {
	*Handler
	groupInvitationService *service.GroupInvitationService
}

func NewGroupInvitationHandler(handler *Handler, groupInvitationService *service.GroupInvitationService) *GroupInvitationHandler {
	return &GroupInvitationHandler{
		Handler:     handler,
		groupInvitationService: groupInvitationService,
	}
}

func (g *GroupInvitationHandler) GetGroupInvitationById(ctx *gin.Context) {

	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	groupInvitation, err := g.groupInvitationService.GetGroupInvitationById(params.Id)
	g.logger.Info("GetGroupInvitationByID", zap.Any("groupInvitation", groupInvitation))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, groupInvitation)
}
