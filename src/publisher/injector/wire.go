//go:build wireinject
// +build wireinject

package publisher

import (
	publisherHandler "gobook/src/publisher/delivery/http"
	publisherRepository "gobook/src/publisher/repository"
	publisherService "gobook/src/publisher/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *publisherHandler.PublisherHandler {
	wire.Build(
		publisherRepository.NewPublisherRepository,
		publisherService.NewPublisherService,
		publisherHandler.NewPublisherHandler,
	)

	return &publisherHandler.PublisherHandler{}
}
