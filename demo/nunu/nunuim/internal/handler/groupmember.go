package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/service"
	"nunuim/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type GroupMemberHandler struct {
	*Handler
	groupMemberService *service.GroupMemberService
}

func NewGroupMemberHandler(handler *Handler, groupMemberService *service.GroupMemberService) *GroupMemberHandler {
	return &GroupMemberHandler{
		Handler:     handler,
		groupMemberService: groupMemberService,
	}
}

func (g *GroupMemberHandler) GetGroupMemberById(ctx *gin.Context) {

	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	groupMember, err := g.groupMemberService.GetGroupMemberById(params.Id)
	g.logger.Info("GetGroupMemberByID", zap.Any("groupMember", groupMember))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, groupMember)
}
