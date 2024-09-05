package handler

import (
	"net/http"

	"github.com/BoruTamena/internal/core/entity"
	models "github.com/BoruTamena/internal/core/models/request"
	"github.com/BoruTamena/internal/core/models/request/params"
	"github.com/BoruTamena/internal/core/port/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type courseHandler struct {
	router  *gin.Engine
	service service.CourseService
	lg      *zap.Logger
}

func NewCourseHandler(router *gin.Engine, sv service.CourseService, logger *zap.Logger) *courseHandler {

	return &courseHandler{
		router:  router,
		service: sv,
		lg:      logger,
	}

}

func (ch courseHandler) InitRouter() {
	// endpoints
}

func (ch courseHandler) CreateCourse(c *gin.Context) {

	var course models.Courses

	if err := c.ShouldBind(&course); err != nil {
		c.Error(err)
	}
	//  creating new courses
	if err := ch.service.CreateNewCourse(c.Request.Context(), course, ch.lg); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "course created successfully "})

}
func (ch courseHandler) AddModule(c *gin.Context) {

	var module models.Modules

	if err := c.ShouldBind(&module); err != nil {

		err = entity.ReqErr.Wrap(err, "Badrequest").WithProperty(entity.Code, http.StatusBadRequest)
		c.Error(err)
		return

	}

	// other code goes here

}

func (ch courseHandler) GetCourseByQuery(c *gin.Context) {

	var param params.CourseParams
	// getting  courses

	if err := c.ShouldBindQuery(&param); err != nil {

		err = entity.AuthErr.Wrap(err, "binding error").WithProperty(entity.Code, 500)
		c.Error(err)
		return
	}

	// getting courses

	courses, err := ch.service.GetCoursesByParams(c.Request.Context(), param, ch.lg)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, courses)

}

func (ch courseHandler) GetPopularCourse(c *gin.Context) {

	// getting popular courses

}
