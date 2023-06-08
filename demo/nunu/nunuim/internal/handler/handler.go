package handler

import (
	"github.com/gin-gonic/gin"
	"nunuim/internal/middleware"
	"nunuim/pkg/log"
	"github.com/sony/sonyflake"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger, sf *sonyflake.Sonyflake) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*middleware.MyCustomClaims).UserId
}
