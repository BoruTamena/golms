package service

import (
	"context"

	models "github.com/BoruTamena/internal/core/models/request"
	"github.com/BoruTamena/internal/core/models/request/params"
	"github.com/BoruTamena/internal/core/models/response"
	"go.uber.org/zap"
)

type CourseService interface {
	CreateNewCourse(context.Context, models.Courses, *zap.Logger) error
	AddCourseModule(context.Context, models.Modules, *zap.Logger) error
	// AddModuleLesson()
	GetCoursesByParams(context.Context, params.CourseParams, *zap.Logger) (response.CourseResponse, error)
}
