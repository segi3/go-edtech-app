//go:build wireinject
// +build wireinject

package class_room

import (
	handler "edtech-app/internal/class_room/delivery/http"
	repository "edtech-app/internal/class_room/repository"
	useCase "edtech-app/internal/class_room/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ClassRoomHandler {
	wire.Build(
		handler.NewClassRoomHandler,
		repository.NewClassRoomRepository,
		useCase.NewClassRoomUseCase,
	)

	return &handler.ClassRoomHandler{}
}
