// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package profile

import (
	"gorm.io/gorm"
	"edtech-app/internal/profile/delivery/http"
	profile2 "edtech-app/internal/profile/usecase"
	"edtech-app/internal/user/repository"
	user2 "edtech-app/internal/user/usecase"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *profile.ProfileHandler {
	userRepository := user.NewUserRepository(db)
	userUseCase := user2.NewUserUseCase(userRepository)
	profileUseCase := profile2.NewProfileUseCase(userUseCase)
	profileHandler := profile.NewProfileHandler(profileUseCase)
	return profileHandler
}
