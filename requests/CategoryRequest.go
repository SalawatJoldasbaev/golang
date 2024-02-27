package requests

type CreateCategoryRequest struct {
	Title string `json:"title" binding:"required"`
}
