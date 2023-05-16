package course

import (
	lessonEntity "edtech-app/internal/lesson/entity"
	productEntity "edtech-app/internal/product/entity"
)

type CourseResponseBody struct {
	Product productEntity.Product `json:"product"`
	Lessons []lessonEntity.Lesson `json:"lesson"`
}
