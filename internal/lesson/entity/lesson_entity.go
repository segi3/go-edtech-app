package lesson

import (
	"database/sql"

	adminEntity "edtech-app/internal/admin/entity"

	"gorm.io/gorm"
)

type Lesson struct {
	ID           int64              `json:"id"`
	Index        *int64             `json:"index" gorm:"->"`
	Title        string             `json:"title"`
	VideoContent *string            `json:"video_content"`
	Description  string             `json:"description"`
	TextContent  string             `json:"text_content"`
	CreatedByID  int64              `json:"created_by" gorm:"column:created_by"`
	CreatedBy    *adminEntity.Admin `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID  *int64             `json:"updated_by"  gorm:"column:updated_by"`
	UpdatedBy    *adminEntity.Admin `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt    sql.NullTime       `json:"created_at"`
	UpdatedAt    sql.NullTime       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt     `json:"deleted_at"`
}
