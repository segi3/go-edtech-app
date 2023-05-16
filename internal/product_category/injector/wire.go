//go:build wireinject
// +build wireinject

package product_category

import (
	handler "edtech-app/internal/product_category/delivery/http"
	repository "edtech-app/internal/product_category/repository"
	usecase "edtech-app/internal/product_category/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ProductCategoryHandler {
	wire.Build(
		handler.NewProductCategoryHandler,
		repository.NewProductCategoryRepository,
		usecase.NewProductCategoryUseCase,
	)

	return &handler.ProductCategoryHandler{}
}
