package repository

import (
	"context"
	"penugasan4/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	GetPostById(ctx context.Context, postId uint64) (entity.Post, error)
	CreatePost(ctx context.Context, post entity.Post) (entity.Post, error)
	LikePostByPostID(ctx context.Context, userID uint64, postID uint64) error
	UpdatePost(ctx context.Context, post entity.Post) error
	DeletePost(ctx context.Context, postID uint64) error
}

type postConnection struct {
	connection     *gorm.DB
	UserRepository UserRepository
}

func NewPostRepository(db *gorm.DB, userRepository UserRepository) PostRepository {
	return &postConnection{
		connection:     db,
		UserRepository: userRepository,
	}
}

func (pc *postConnection) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	var posts []entity.Post
	if tx := pc.connection.Preload("User").Preload("Likes").Preload("Comments").Find(&posts).Error; tx != nil {
		return nil, tx
	}
	return posts, nil
}

func (pc *postConnection) GetPostById(ctx context.Context, postId uint64) (entity.Post, error) {
	var post entity.Post

	if tx := pc.connection.Preload("User").Preload("Likes").Preload("Comments").Where("id = ?", postId).Take(&post).Error; tx != nil {
		return entity.Post{}, tx
	}
	return post, nil
}

func (pc *postConnection) CreatePost(ctx context.Context, post entity.Post) (entity.Post, error) {
	if tx := pc.connection.Create(&post).Error; tx != nil {
		return entity.Post{}, tx
	}

	return post, nil
}

func (pc *postConnection) LikePostByPostID(ctx context.Context, userID uint64, postID uint64) error {
	var like entity.Like
	if tx := pc.connection.Where("user_id = ? AND post_id = ?", userID, postID).Find(&like).Error; tx != nil {
		return tx
	}

	like = entity.Like{
		PostID: postID,
		UserID: userID,
	}

	if tx := pc.connection.Create(&like).Error; tx != nil {
		return tx
	}

	var post entity.Post
	if tx := pc.connection.Where("id = ?", postID).Find(&post).Error; tx != nil {
		return tx
	}

	pc.UpdatePost(ctx, post)
	return nil
}

func (pc *postConnection) UpdatePost(ctx context.Context, post entity.Post) error {
	if tx := pc.connection.Save(&post).Error; tx != nil {
		return tx
	}
	return nil
}

func (pc *postConnection) DeletePost(ctx context.Context, postID uint64) error {
	if tx := pc.connection.Delete(&entity.Post{}, "id = ?", &postID).Error; tx != nil {
		return tx
	}
	return nil
}
