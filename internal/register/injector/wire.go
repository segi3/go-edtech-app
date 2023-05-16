//go:build wireinject
// +build wireinject

package register

// mail := mail.NewMail()
// 	userRepository := userRepository.NewUserRepository(db)
// 	userUseCase := userUseCase.NewUserUseCase(userRepository)
// 	registerUseCase := registerUseCase.NewRegisterUseCase(userUseCase, mail)
// 	registerHandler.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)

import (
	handler "edtech-app/internal/register/delivery/http"
	useCase "edtech-app/internal/register/usecase"
	userRepository "edtech-app/internal/user/repository"
	userUseCase "edtech-app/internal/user/usecase"
	mail "edtech-app/pkg/mail/sendgrid"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.RegisterHandler {
	wire.Build(
		handler.NewRegisterHandler,
		useCase.NewRegisterUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
		mail.NewMail,
	)

	return &handler.RegisterHandler{}
}
