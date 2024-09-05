package service

import (
	"github.com/BoruTamena/internal/core/entity"
	"github.com/BoruTamena/internal/core/models/request"
	"github.com/BoruTamena/internal/core/port/repository"
	"github.com/BoruTamena/internal/core/port/service"
	"go.uber.org/zap"
)

type courseServiceImpl struct {
	rep repository.CourseRepository
}

func NewCourseRepository(courseRep repository.CourseRepository) service.CourseService {

	return &courseServiceImpl{
		rep: courseRep,
	}

}

// creating new courses
func (cs courseServiceImpl) CreateNewCourse(course request.Courses, lg *zap.Logger) error {

	if err := course.Validate(); err != nil {

		err = entity.ValErr.Wrap(err, "validation error").WithProperty(entity.Code, 500)

		lg.Error("error",
			zap.Error(err),
			zap.String("message", err.Error()),
		)

		return err
	}

	// other code goes here ...

	return nil

}
