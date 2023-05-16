package course

type CourseBindingRequestBody struct {
	ProductID int64   `json:"product_id"`
	LessonIDs []int64 `json:"lesson_ids"`
	CreatedBy int64   `json:"created_by"`
	UpdatedBy int64   `json:"updated_by"`
}
