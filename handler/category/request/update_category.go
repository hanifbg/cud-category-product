package request

type UpdateCategoryRequest struct {
	Name     string `json:"name"`
	IsActive int    `json:"is_active"`
}
