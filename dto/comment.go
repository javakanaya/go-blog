package dto

type CreateCommentDTO struct {
	ID      uint64 `json:"id"`
	PostID  uint64 `json:"post_id" binding:"required"`
	UserID  uint64 `json:"user_id"`
	Comment string `json:"comment" binding:"required"`
}

type UpdateCommentDTO struct {
	ID      uint64 `json:"id"`
	PostID  uint64 `json:"post_id" binding:"required"`
	UserID  uint64 `json:"user_id"`
	Comment string `json:"comment" binding:"required"`
}
