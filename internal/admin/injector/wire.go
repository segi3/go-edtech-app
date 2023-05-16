//go:build wireinject
// +build wireinject

package admin

import (
	handler "edtech-app/internal/admin/delivery/http"
	repository "edtech-app/internal/admin/repository"
	usecase "edtech-app/internal/admin/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.AdminHandler {
	wire.Build(
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		handler.NewAdminHandler,
	)

	return &handler.AdminHandler{}
}
