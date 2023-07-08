//go:build wireinject
// +build wireinject

package register

import (
	registerHandler "gobook/src/register/delivery/http"
	registerService "gobook/src/register/service"

	"github.com/google/wire"
	"gorm.io/gorm"

	userRepository "gobook/src/user/repository"
	userService "gobook/src/user/service"

	mail "gobook/pkg/mail/gomail"
)

func InitializeService(db *gorm.DB) *registerHandler.RegisterHandler {
	wire.Build(
		registerHandler.NewRegisterHandler,
		userRepository.NewUserRepository,
		registerService.NewRegisterService,
		userService.NewUserService,
		mail.NewSmtpMail,
	)

	return &registerHandler.RegisterHandler{}
}
