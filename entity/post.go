package entity

import "time"

type Post struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      string    `json:"tags"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`

	UserID   uint64    `json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"user"`
	Likes    []Like    `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"likes,omitempy"`
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Comment,omitempy"`
}
