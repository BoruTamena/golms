package service

import (
	"context"
	"net/http"

	"github.com/BoruTamena/internal/core/entity"
	models "github.com/BoruTamena/internal/core/models/request"
	"github.com/BoruTamena/internal/core/models/request/params"
	"github.com/BoruTamena/internal/core/models/response"
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
func (cs courseServiceImpl) CreateNewCourse(ctx context.Context, course models.Courses, lg *zap.Logger) error {

	if err := course.Validate(); err != nil {

		err = entity.ValErr.Wrap(err, "validation error").WithProperty(entity.Code, http.StatusBadRequest)

		lg.Error("error",
			zap.Error(err),
			zap.String("message", err.Error()),
		)

		return err
	}

	// other code goes here ...

	return nil

}

func (cs courseServiceImpl) AddCourseModule(c context.Context, module models.Modules, lg *zap.Logger) error {

	if err := module.Validate(); err != nil {

		err = entity.ValErr.Wrap(err, "Adding course module validation error").WithProperty(entity.Code, http.StatusBadRequest)
		lg.Error("error",
			zap.Error(err),
			zap.String("message", err.Error()),
		)

		return err
	}

	// other code goes here ...

}

func (cs courseServiceImpl) GetCoursesByParams(ctx context.Context, query_param params.CourseParams, lg *zap.Logger) (response.CourseResponse, error) {

	if err := query_param.Validate(); err != nil {

		err = entity.ValErr.Wrap(err, "query parameters validation error ").WithProperty(entity.Code, http.StatusBadRequest)
		lg.Error("error",
			zap.Error(err),
			zap.String("message", err.Error()),
		)
		return response.CourseResponse{
			Status: http.StatusBadRequest,
		}, err

	}

	// other code goes here ...
}
