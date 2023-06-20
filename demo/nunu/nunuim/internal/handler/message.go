package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/service"
	"nunuim/pkg/helper/resp"
	"go.uber.org/zap"
	"net/http"
)

type MessageHandler struct {
	*Handler
	messageService *service.MessageService
}

func NewMessageHandler(handler *Handler, messageService *service.MessageService) *MessageHandler {
	return &MessageHandler{
		Handler:     handler,
		messageService: messageService,
	}
}

func (m *MessageHandler) GetMessageById(ctx *gin.Context) {

	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	message, err := m.messageService.GetMessageById(params.Id)
	m.logger.Info("GetMessageByID", zap.Any("message", message))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, message)
}
