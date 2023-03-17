package service

import (
	"context"
	"penugasan4/dto"
	"penugasan4/entity"
	"penugasan4/repository"
	"penugasan4/utils"

	"github.com/mashingan/smapping"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDTO dto.UserRegister) (entity.User, error)
	CheckUserByEmail(ctx context.Context, email string) bool
	VerifyUser(ctx context.Context, email string, passsword string) (bool, error)
	GetUserByID(ctx context.Context, userId uint64) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateUser(ctx context.Context, userDTO dto.UserUpdate) error
	DeleteUser(ctx context.Context, userID uint64) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) RegisterUser(ctx context.Context, userDTO dto.UserRegister) (entity.User, error) {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(userDTO))
	if err != nil {
		return user, err
	}
	user.Role = "user"
	return us.userRepository.RegisterUser(ctx, user)
}

func (us *userService) CheckUserByEmail(ctx context.Context, email string) bool {
	res, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false
	}
	if res.Email == "" {
		return false
	}
	return true
}

func (us *userService) VerifyUser(ctx context.Context, email string, password string) (bool, error) {
	res, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	checkPassword, err := utils.PasswordCompare(res.Password, []byte(password))
	if err != nil {
		return false, err
	}

	if res.Email == email && checkPassword {
		return true, nil
	}
	return false, nil
}

func (us *userService) GetUserByID(ctx context.Context, userID uint64) (entity.User, error) {
	return us.userRepository.GetUserByID(ctx, userID)
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return us.userRepository.GetUserByEmail(ctx, email)
}

func (us *userService) UpdateUser(ctx context.Context, userDTO dto.UserUpdate) error {
	user := entity.User{}
	if err := smapping.FillStruct(&user, smapping.MapFields(userDTO)); err != nil {
		return nil
	}
	return us.userRepository.UpdateUser(ctx, user)
}

func (us *userService) DeleteUser(ctx context.Context, userID uint64) error {
	return us.userRepository.DeleteUser(ctx, userID)
}
