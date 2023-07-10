//go:build wireinject
// +build wireinject

package book

import (
	bookHandler "gobook/src/book/delivery/http"
	bookRepository "gobook/src/book/repository"
	bookService "gobook/src/book/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *bookHandler.BookHandler {
	wire.Build(
		bookRepository.NewBookRepository,
		bookService.NewBookService,
		bookHandler.NewBookHandler,
	)

	return &bookHandler.BookHandler{}
}
