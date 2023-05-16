package course

import (
	"database/sql"

	adminEntity "edtech-app/internal/admin/entity"
	lessonEntity "edtech-app/internal/lesson/entity"
	productEntity "edtech-app/internal/product/entity"

	"gorm.io/gorm"
)

type Course struct {
	ID          int64                  `json:"id"`
	Lesson      *lessonEntity.Lesson   `json:"leson" gorm:"foreignKey:LessonID;references:ID"`
	LessonID    int64                  `json:"lesson_id"`
	Index       *int64                 `json:"index"`
	Product     *productEntity.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID   int64                  `json:"product_id"`
	CreatedByID *int64                 `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *adminEntity.Admin     `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64                 `json:"updated_by"  gorm:"column:updated_by"`
	UpdatedBy   *adminEntity.Admin     `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   sql.NullTime           `json:"created_at"`
	UpdatedAt   sql.NullTime           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt         `json:"deleted_at"`
}
