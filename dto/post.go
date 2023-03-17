package dto

type PostCreateDTO struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint64 `json:"user_id"`
}

type PostUpdateDTO struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint64 `json:"user_id"`
}
