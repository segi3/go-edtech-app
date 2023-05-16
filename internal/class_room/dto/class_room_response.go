package class_room

import (
	"database/sql"

	classRoomEntity "edtech-app/internal/class_room/entity"
	productEntity "edtech-app/internal/product/entity"
	userEntity "edtech-app/internal/user/entity"

	"gorm.io/gorm"
)

type ClassRoomResponseBody struct {
	ID        int64                  `json:"id"`
	User      *userEntity.User       `json:"user"`
	Product   *productEntity.Product `json:"product"`
	CreatedBy *userEntity.User       `json:"created_by"`
	UpdatedBy *userEntity.User       `json:"updated_by"`
	CreatedAt sql.NullTime           `json:"created_at"`
	UpdatedAt sql.NullTime           `json:"updated_at"`
	DeletedAt gorm.DeletedAt         `json:"deleted_at"`
}

func CreateClassRoomResponse(entity classRoomEntity.ClassRoom) ClassRoomResponseBody {
	return ClassRoomResponseBody{
		ID:        entity.ID,
		User:      entity.User,
		Product:   entity.Product,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

type ClassRoomListResponse []ClassRoomResponseBody

func CreateClassRoomListResponse(entity []classRoomEntity.ClassRoom) ClassRoomListResponse {
	classRoomResp := ClassRoomListResponse{}

	for _, classRoom := range entity {
		classRoom.Product.VideoURL = classRoom.Product.Video

		classRoomResponse := CreateClassRoomResponse(classRoom)
		classRoomResp = append(classRoomResp, classRoomResponse)
	}

	return classRoomResp
}
