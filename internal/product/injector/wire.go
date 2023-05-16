//go:build wireinject
// +build wireinject

package product

import (
	handler "edtech-app/internal/product/delivery/http"
	repository "edtech-app/internal/product/repository"
	usecase "edtech-app/internal/product/usecase"
	fileUpload "edtech-app/pkg/fileupload/cloudinary"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitiliazedService(db *gorm.DB) *handler.ProductHandler {
	wire.Build(
		handler.NewProductHandler,
		usecase.NewProductUseCase,
		repository.NewProductRepository,
		fileUpload.NewFileUpload,
	)

	return &handler.ProductHandler{}
}
