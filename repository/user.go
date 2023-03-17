package repository

import (
	"context"
	"penugasan4/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserByID(ctx context.Context, userID uint64) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) error
	DeleteUser(ctx context.Context, UserID uint64) error
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) RegisterUser(ctx context.Context, user entity.User) (entity.User, error) {
	if tx := db.connection.Create(&user).Error; tx != nil {
		return entity.User{}, tx
	}
	return user, nil
}

func (db *userConnection) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if tx := db.connection.Where("email = ?", email).Take(&user).Error; tx != nil {
		return entity.User{}, tx
	}
	return user, nil
}

func (db *userConnection) GetUserByID(ctx context.Context, userID uint64) (entity.User, error) {
	var user entity.User
	if tx := db.connection.Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "title")
	}).Preload("Likes").Preload("Comments").Where("id = ?", userID).Take(&user).Error; tx != nil {
		return entity.User{}, tx
	}
	return user, nil
}

func (db *userConnection) UpdateUser(ctx context.Context, user entity.User) error {
	if tx := db.connection.Updates(&user).Error; tx != nil {
		return tx
	}
	return nil
}
func (db *userConnection) DeleteUser(ctx context.Context, userID uint64) error {
	if tx := db.connection.Delete(&entity.User{}, &userID).Error; tx != nil {
		return tx
	}
	return nil
}
