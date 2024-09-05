package service

import (
	models "github.com/BoruTamena/internal/core/models/request"
	"go.uber.org/zap"
)

type UserService interface {
	Register(models.User, *zap.Logger) error
}
