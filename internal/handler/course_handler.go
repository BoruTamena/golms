package handler

import (
	"net/http"

	"github.com/BoruTamena/internal/core/models/request"
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

	var course request.Courses

	if err := c.ShouldBind(&course); err != nil {
		c.Error(err)
	}
	//  creating new courses
	if err := ch.service.CreateNewCourse(course, ch.lg); err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "course created successfully "})

}

func (ch courseHandler) GetCourse(c *gin.Context) {

	// getting all courses

}

func (ch courseHandler) GetPopularCourse(c *gin.Context) {

	// getting popular courses

}

func (ch courseHandler) GetCourseByCategory(c *gin.Context) {

	// getting course of specified category

}
