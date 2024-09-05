package service

import (
	"github.com/BoruTamena/internal/core/entity"
	models "github.com/BoruTamena/internal/core/models/request"
	"github.com/BoruTamena/internal/core/port/repository"
	"github.com/BoruTamena/internal/core/port/service"
	"go.uber.org/zap"
)

type userServiceImpl struct {
	rep repository.UserRepository
}

func NewUserService(rep repository.UserRepository) service.UserService {

	return &userServiceImpl{
		rep: rep,
	}

}

func (us userServiceImpl) Register(user models.User, lg *zap.Logger) error {

	if err := user.Validate(); err != nil {

		err = entity.ValErr.Wrap(err, "validation error").WithProperty(entity.Code, 500)
		// logging
		lg.Error("error",
			zap.String("message", "djsdkl"),
			zap.Error(err),
		)
		return err

	}

	// other code goes here

	return nil
}
