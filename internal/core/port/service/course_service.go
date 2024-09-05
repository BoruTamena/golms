package service

import (
	"github.com/BoruTamena/internal/core/models/request"
	"go.uber.org/zap"
)

type CourseService interface {
	CreateNewCourse(request.Courses, *zap.Logger) error
}
