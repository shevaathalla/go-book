// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rental

import (
	"gobook/pkg/fileupload/cloudinary"
	"gobook/src/rental/delivery/http"
	rental2 "gobook/src/rental/repository"
	rental3 "gobook/src/rental/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeService(db *gorm.DB) *rental.RentalHandler {
	rentalRepository := rental2.NewRentalRepository(db)
	fileUpload := fileupload.NewFileUpload()
	rentalService := rental3.NewRentalService(rentalRepository, fileUpload)
	rentalHandler := rental.NewRentalHandler(rentalService)
	return rentalHandler
}
