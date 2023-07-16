//go:build wireinject
// +build wireinject

package rental

import (
	rentalHandler "gobook/src/rental/delivery/http"
	rentalRepository "gobook/src/rental/repository"
	rentalService "gobook/src/rental/service"

	fileupload "gobook/pkg/fileupload/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *rentalHandler.RentalHandler {
	wire.Build(
		rentalRepository.NewRentalRepository,
		rentalService.NewRentalService,
		rentalHandler.NewRentalHandler,
		fileupload.NewFileUpload,
	)

	return &rentalHandler.RentalHandler{}
}
