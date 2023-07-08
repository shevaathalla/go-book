//go:build wireinject
// +build wireinject

package user

import (
	userHandler "gobook/src/user/delivery/http"
	userRepository "gobook/src/user/repository"
	userService "gobook/src/user/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *userHandler.UserHandler {
	wire.Build(
		userRepository.NewUserRepository,
		userService.NewUserService,
		userHandler.NewUserHandler,
	)

	return &userHandler.UserHandler{}
}
