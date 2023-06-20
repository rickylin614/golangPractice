package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/service"
	"nunuim/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type GroupRequestHandler struct {
	*Handler
	groupRequestService *service.GroupRequestService
}

func NewGroupRequestHandler(handler *Handler, groupRequestService *service.GroupRequestService) *GroupRequestHandler {
	return &GroupRequestHandler{
		Handler:     handler,
		groupRequestService: groupRequestService,
	}
}

func (g *GroupRequestHandler) GetGroupRequestById(ctx *gin.Context) {

	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	groupRequest, err := g.groupRequestService.GetGroupRequestById(params.Id)
	g.logger.Info("GetGroupRequestByID", zap.Any("groupRequest", groupRequest))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, groupRequest)
}
