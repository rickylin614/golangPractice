package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/service"
	"nunuim/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type GroupHandler struct {
	*Handler
	groupService *service.GroupService
}

func NewGroupHandler(handler *Handler, groupService *service.GroupService) *GroupHandler {
	return &GroupHandler{
		Handler:     handler,
		groupService: groupService,
	}
}

func (g *GroupHandler) GetGroupById(ctx *gin.Context) {

	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	group, err := g.groupService.GetGroupById(params.Id)
	g.logger.Info("GetGroupByID", zap.Any("group", group))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, group)
}
