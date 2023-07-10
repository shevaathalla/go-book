//go:build wireinject
// +build wireinject

package author

import (
	authorHandler "gobook/src/author/delivery/http"
	authorRepository "gobook/src/author/repository"
	authorService "gobook/src/author/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *authorHandler.AuthorHandler {
	wire.Build(
		authorRepository.NewAuthorRepository,
		authorService.NewAuthorService,
		authorHandler.NewAuthorHandler,
	)

	return &authorHandler.AuthorHandler{}
}
