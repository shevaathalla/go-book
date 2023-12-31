// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package publisher

import (
	"gobook/src/publisher/delivery/http"
	"gobook/src/publisher/repository"
	"gobook/src/publisher/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeService(db *gorm.DB) *publisher.PublisherHandler {
	publisherRepository := repository.NewPublisherRepository(db)
	publisherService := service.NewPublisherService(publisherRepository)
	publisherHandler := publisher.NewPublisherHandler(publisherService)
	return publisherHandler
}
