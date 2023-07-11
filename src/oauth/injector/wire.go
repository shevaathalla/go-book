//go:build wireinject
// +build wireinject

package oauth

import (
	oauthHandler "gobook/src/oauth/delivery/http"
	oauthService "gobook/src/oauth/service"
	userRepository "gobook/src/user/repository"
	userService "gobook/src/user/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *oauthHandler.OauthHandler {
	wire.Build(
		oauthService.NewOauthService,
		oauthHandler.NewOauthHandler,
		userService.NewUserService,
		userRepository.NewUserRepository,
	)
	return &oauthHandler.OauthHandler{}
}
