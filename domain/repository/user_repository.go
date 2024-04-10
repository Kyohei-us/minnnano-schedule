package repository

import (
	"context"
	"minnnano-schedule/domain/model"
)

type IUserRepository interface {
	SelectUsers(ctx context.Context) ([]*model.User, error)
	SelectUserById(ctx context.Context, userID int) (*model.User, error)
}
