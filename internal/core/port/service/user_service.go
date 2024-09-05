package service

import (
	"github.com/BoruTamena/internal/core/models/request"
	"go.uber.org/zap"
)

type UserService interface {
	Register(request.User, *zap.Logger) error
}
