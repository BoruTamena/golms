package handler

import (
	"net/http"

	"github.com/BoruTamena/internal/core/entity"
	"github.com/BoruTamena/internal/core/models/request"
	"github.com/BoruTamena/internal/core/port/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userHandler struct {
	router  *gin.Engine
	service service.UserService
	lg      *zap.Logger
}

func NewHandler(r *gin.Engine, s service.UserService, logger *zap.Logger) *userHandler {

	return &userHandler{
		router:  r,
		service: s,
		lg:      logger,
	}
}

func (engine *userHandler) InitRouter() {

	// endpoints
	engine.router.POST("/user", engine.SignUP)

}

// defining endpoints

func (handler *userHandler) SignUP(c *gin.Context) {

	var user request.User
	// binding user data
	err := c.ShouldBind(&user)

	if err != nil {

		err := entity.ValErr.Wrap(err, "can't bind user data for some reason").WithProperty(entity.Code, 500)
		handler.lg.Error("error",
			zap.String("message", err.Error()),
		)
		c.Error(err)
	}

	// register user
	if err = handler.service.Register(user, handler.lg); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully "})

}
