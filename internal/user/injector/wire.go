//go:build wireinject
// +build wireinject

package user

import (
	handler "edtech-app/internal/user/delivery/http"
	repository "edtech-app/internal/user/repository"
	useCase "edtech-app/internal/user/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.UserHandler {
	wire.Build(
		handler.NewUserHandler,
		repository.NewUserRepository,
		useCase.NewUserUseCase,
	)

	return &handler.UserHandler{}
}
