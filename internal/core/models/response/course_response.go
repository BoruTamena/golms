package response

import models "github.com/BoruTamena/internal/core/models/request"

type CourseResponse struct {
	Status int `json:"status"`
	Data   []models.Courses
}
