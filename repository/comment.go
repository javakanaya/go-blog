package repository

import (
	"context"
	"penugasan4/entity"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	UpdateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	DeleteComment(ctx context.Context, userID uint64) error
}

type commentRepository struct {
	connection     *gorm.DB
	userRepository UserRepository
}

func NewCommentRepository(db *gorm.DB, ur UserRepository) CommentRepository {
	return &commentRepository{
		connection:     db,
		userRepository: ur,
	}
}

func (cr *commentRepository) CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	if tx := cr.connection.Create(&comment).Error; tx != nil {
		return entity.Comment{}, tx
	}
	return comment, nil
}

func (cr *commentRepository) UpdateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	if tx := cr.connection.Save(&comment).Error; tx != nil {
		return entity.Comment{}, tx
	}
	return comment, nil
}

func (cr *commentRepository) DeleteComment(ctx context.Context, commentID uint64) error {
	if tx := cr.connection.Delete(&entity.Comment{}, &commentID).Error; tx != nil {
		return tx
	}
	return nil
}
