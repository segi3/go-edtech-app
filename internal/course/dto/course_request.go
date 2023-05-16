package course

type CourseBindingRequestBody struct {
	ProductID   int64   `json:"product_id" binding:"required"`
	LessonIDs   []int64 `json:"lesson_ids" binding:"required"`
	LessonIndex []int64 `json:"lesson_index" binding:"required"`
	CreatedBy   int64   `json:"created_by"`
	UpdatedBy   int64   `json:"updated_by"`
}
