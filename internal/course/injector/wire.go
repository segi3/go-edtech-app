//go:build wireinject
// +build wireinject

package product

import (
	handler "edtech-app/internal/course/delivery/http"
	repository "edtech-app/internal/course/repository"
	usecase "edtech-app/internal/course/usecase"
	fileUpload "edtech-app/pkg/fileupload/cloudinary"

	lessonRepository "edtech-app/internal/lesson/repository"
	lessonUseCase "edtech-app/internal/lesson/usecase"
	productRepository "edtech-app/internal/product/repository"
	productUseCase "edtech-app/internal/product/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitiliazedService(db *gorm.DB) *handler.CourseHandler {
	wire.Build(
		handler.NewCourseHandler,
		usecase.NewCourseUseCase,
		repository.NewCourseRepository,
		fileUpload.NewFileUpload,
		lessonUseCase.NewLessonUseCase,
		productUseCase.NewProductUseCase,
		lessonRepository.NewLessonRepository,
		productRepository.NewProductRepository,
	)

	return &handler.CourseHandler{}
}
