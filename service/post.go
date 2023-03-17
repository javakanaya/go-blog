package service

import (
	"penugasan4/dto"
	"penugasan4/entity"
	"penugasan4/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/net/context"
)

type PostService interface {
	CreatePost(ctx context.Context, postDTO dto.PostCreateDTO) (entity.Post, error)
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	GetPostById(ctx context.Context, postID uint64) (entity.Post, error)
	UpdatePost(ctx context.Context, postDTO dto.PostUpdateDTO) error
	LikePostByPostID(ctx context.Context, userID uint64, postID uint64) error
	DeletePost(ctx context.Context, postID uint64) error
}

type postService struct {
	postRepository repository.PostRepository
	userRepository repository.UserRepository
}

func NewPostService(pr repository.PostRepository, ur repository.UserRepository) PostService {
	return &postService{
		postRepository: pr,
		userRepository: ur,
	}
}

func (ps *postService) CreatePost(ctx context.Context, postDTO dto.PostCreateDTO) (entity.Post, error) {
	var post entity.Post
	if err := smapping.FillStruct(&post, smapping.MapFields(postDTO)); err != nil {
		return post, err
	}

	return ps.postRepository.CreatePost(ctx, post)
}

func (ps *postService) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	return ps.postRepository.GetAllPosts(ctx)
}

func (ps *postService) GetPostById(ctx context.Context, postID uint64) (entity.Post, error) {
	return ps.postRepository.GetPostById(ctx, postID)
}

func (ps *postService) UpdatePost(ctx context.Context, postDTO dto.PostUpdateDTO) error {
	var post entity.Post
	if err := smapping.FillStruct(&post, smapping.MapFields(postDTO)); err != nil {
		return err
	}
	return ps.postRepository.UpdatePost(ctx, post)
}

func (ps *postService) LikePostByPostID(ctx context.Context, userID uint64, postID uint64) error {
	return ps.postRepository.LikePostByPostID(ctx, userID, postID)
}

func (ps *postService) DeletePost(ctx context.Context, postID uint64) error {
	return ps.postRepository.DeletePost(ctx, postID)
}
