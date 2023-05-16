//go:build wireinject
// +build wireinject

package oauth

import (
	adminRepository "edtech-app/internal/admin/repository"
	adminUseCase "edtech-app/internal/admin/usecase"
	oauthHandler "edtech-app/internal/oauth/delivery/http"
	oauthRepository "edtech-app/internal/oauth/repository"
	oauthUseCase "edtech-app/internal/oauth/usecase"
	userRepository "edtech-app/internal/user/repository"
	userUseCase "edtech-app/internal/user/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *oauthHandler.OauthHandler {
	wire.Build(
		oauthHandler.NewOauthHandler,
		oauthRepository.NewOauthClientRepository,
		oauthRepository.NewOauthAccessTokenRepository,
		oauthRepository.NewOauthRefreshTokenRepository,
		oauthUseCase.NewOauthUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
		adminRepository.NewAdminRepository,
		adminUseCase.NewAdminUseCase,
	)

	return &oauthHandler.OauthHandler{}

}
