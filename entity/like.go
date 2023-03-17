package entity

type Like struct {
	ID     uint64 `gorm:"PrimaryKey" json:"id"`
	PostID uint64 `json:"post_id"`
	UserID uint64 `json:"user_id"`
}
