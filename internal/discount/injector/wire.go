//go:build wireinject
// +build wireinject

package discount

import (
	handler "edtech-app/internal/discount/delivery/http"
	repository "edtech-app/internal/discount/repository"
	useCase "edtech-app/internal/discount/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.DiscountHandler {
	wire.Build(
		handler.NewDiscountHandler,
		repository.NewDiscountRepository,
		useCase.NewDiscountUseCase,
	)

	return &handler.DiscountHandler{}
}
