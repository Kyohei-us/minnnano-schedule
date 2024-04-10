package postgresql

import (
	"context"
	"minnnano-schedule/domain/model"
	"minnnano-schedule/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) SelectUsers(ctx context.Context) ([]*model.User, error) {
	// concrete DB operation
	var users []*model.User
	result := ur.DB.Find(&users)
	return users, result.Error
}

func (ur *userRepository) SelectUserById(ctx context.Context, userID int) (*model.User, error) {
	// concrete DB operation
	var user model.User
	result := ur.DB.First(&user, "Id = ?", userID)
	return &user, result.Error
}
