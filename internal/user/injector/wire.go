//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/user/delivery/http"
	repository "online-course/internal/user/repository"
	useCase "online-course/internal/user/usecase"
)

func InitializedService(db *gorm.DB) *handler.UserHandler {
	wire.Build(
		handler.NewUserHandler,
		repository.NewUserRepository,
		useCase.NewUserUseCase,
	)

	return &handler.UserHandler{}
}
