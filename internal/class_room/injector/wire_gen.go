// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package class_room

import (
	"gorm.io/gorm"
	"edtech-app/internal/class_room/delivery/http"
	class_room2 "edtech-app/internal/class_room/repository"
	class_room3 "edtech-app/internal/class_room/usecase"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *class_room.ClassRoomHandler {
	classRoomRepository := class_room2.NewClassRoomRepository(db)
	classRoomUseCase := class_room3.NewClassRoomUseCase(classRoomRepository)
	classRoomHandler := class_room.NewClassRoomHandler(classRoomUseCase)
	return classRoomHandler
}
