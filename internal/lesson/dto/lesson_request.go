package lesson

import "mime/multipart"

type ProductRequestBody struct {
	Title        string                `form:"title" binding:"required"`
	Description  string                `form:"description" binding:"required"`
	VideoContent *multipart.FileHeader `form:"video_content"`
	TextContent  string                `form:"text_content" binding:"required"`
	CreatedBy    int64                 `form:"created_by"`
	UpdatedBy    int64                 `form:"updated_by"`
}
