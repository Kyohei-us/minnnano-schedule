package service

import (
	"context"
	"minnnano-schedule/domain/model"
	"minnnano-schedule/domain/repository"
)

type IUserService interface {
	FindUsers(ctx context.Context) ([]*model.User, error)
	FindUserById(ctx context.Context, id int) (*model.User, error)
	AddUser(ctx context.Context, userName string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}

// インターフェースを満たすstruct
type userService struct {
	repo repository.IUserRepository
}

// NewXXXX(コンストラクタ)については後ほど説明する
func NewUserService(sr repository.IUserRepository) IUserService {
	return &userService{
		repo: sr,
	}
}

func (ss *userService) FindUsers(ctx context.Context) ([]*model.User, error) {
	return ss.repo.SelectUsers(ctx)
}

func (ss *userService) FindUserById(ctx context.Context, id int) (*model.User, error) {
	return ss.repo.SelectUserById(ctx, id)
}

func (ss *userService) AddUser(ctx context.Context, userName string) (*model.User, error) {
	return ss.repo.InsertUser(ctx, userName)
}

func (ss *userService) DeleteUser(ctx context.Context, id int) (int, error) {
	return ss.repo.DeleteUser(ctx, id)
}
