//go:build wireinject
// +build wireinject

package Cart

import (
	handler "edtech-app/internal/cart/delivery/http"
	repository "edtech-app/internal/cart/repository"
	usecase "edtech-app/internal/cart/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitiliazedService(db *gorm.DB) *handler.CartHandler {
	wire.Build(
		handler.NewCartHandler,
		repository.NewCartRepository,
		usecase.NewCartUseCase,
	)

	return &handler.CartHandler{}
}
