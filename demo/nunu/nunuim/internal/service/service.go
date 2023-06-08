package service

import (
	"nunuim/internal/middleware"
	"nunuim/pkg/log"
	"github.com/sony/sonyflake"
)

type Service struct {
	logger    *log.Logger
	sonyflake *sonyflake.Sonyflake
	jwt       *middleware.JWT
}

func NewService(logger *log.Logger, sonyflake *sonyflake.Sonyflake, jwt *middleware.JWT) *Service {
	return &Service{
		logger:    logger,
		sonyflake: sonyflake,
		jwt:       jwt,
	}
}
