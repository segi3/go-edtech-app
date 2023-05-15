// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package register

import (
	"gorm.io/gorm"
	"online-course/internal/register/delivery/http"
	register2 "online-course/internal/register/usecase"
	"online-course/internal/user/repository"
	user2 "online-course/internal/user/usecase"
	"online-course/pkg/mail/sendgrid"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *register.RegisterHandler {
	userRepository := user.NewUserRepository(db)
	userUseCase := user2.NewUserUseCase(userRepository)
	mailMail := mail.NewMail()
	registerUseCase := register2.NewRegisterUseCase(userUseCase, mailMail)
	registerHandler := register.NewRegisterHandler(registerUseCase)
	return registerHandler
}