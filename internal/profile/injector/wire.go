//go:build wireinject
// +build wireinject

package profile

import (
	handler "edtech-app/internal/profile/delivery/http"
	useCase "edtech-app/internal/profile/usecase"
	userRepository "edtech-app/internal/user/repository"
	userUseCase "edtech-app/internal/user/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ProfileHandler {
	wire.Build(
		handler.NewProfileHandler,
		useCase.NewProfileUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &handler.ProfileHandler{}
}
