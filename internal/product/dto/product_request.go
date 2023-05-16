package product

import "mime/multipart"

type ProductRequestBody struct {
	ProductCategoryID int64                 `form:"product_category_id" binding:"required"`
	Title             string                `form:"title" binding:"required"`
	Thumbnail         *multipart.FileHeader `form:"thumbnail"`
	Video             *multipart.FileHeader `form:"video"`
	Description       string                `form:"description" binding:"required"`
	Price             int64                 `form:"price" binding:"required"`
	Duration          int64                 `form:"duration" binding:"required"`
	CreatedBy         int64                 `form:"created_by"`
	UpdatedBy         int64                 `form:"updated_by"`
}
