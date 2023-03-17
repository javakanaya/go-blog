package service

import (
	"context"
	"penugasan4/dto"
	"penugasan4/entity"
	"penugasan4/repository"

	"github.com/mashingan/smapping"
)

type CommentService interface {
	CreateComment(ctx context.Context, commentDTO dto.CreateCommentDTO) (entity.Comment, error)
	UpdateComment(ctx context.Context, commentDTO dto.CreateCommentDTO) (entity.Comment, error)
	DeleteComment(ctx context.Context, userID uint64) error
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(cr repository.CommentRepository) CommentService {
	return &commentService{
		commentRepository: cr,
	}
}

func (cs *commentService) CreateComment(ctx context.Context, commentDTO dto.CreateCommentDTO) (entity.Comment, error) {
	var comment entity.Comment
	if err := smapping.FillStruct(&comment, smapping.MapFields(commentDTO)); err != nil {
		return entity.Comment{}, err
	}
	return cs.commentRepository.CreateComment(ctx, comment)
}

func (cs *commentService) UpdateComment(ctx context.Context, commentDTO dto.CreateCommentDTO) (entity.Comment, error) {
	var comment entity.Comment
	if err := smapping.FillStruct(&comment, smapping.MapFields(commentDTO)); err != nil {
		return entity.Comment{}, err
	}
	return cs.commentRepository.UpdateComment(ctx, comment)
}

func (cs *commentService) DeleteComment(ctx context.Context, userID uint64) error {
	return cs.commentRepository.DeleteComment(ctx, userID)
}
