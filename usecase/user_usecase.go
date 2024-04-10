package usecase

import (
	"context"
	"minnnano-schedule/domain/service"
	"minnnano-schedule/usecase/model"
)

type IUserUsecase interface {
	FindUsers(ctx context.Context) ([]*model.User, error)
	FindUserById(ctx context.Context, id int) (*model.User, error) // 追加
}

// インターフェースを満たすstruct
type userUsecase struct {
	svc service.IUserService
}

func NewUserUsecase(ss service.IUserService) IUserUsecase {
	return &userUsecase{
		svc: ss,
	}
}

func (uu *userUsecase) FindUsers(ctx context.Context) ([]*model.User, error) {
	msSlice, err := uu.svc.FindUsers(ctx)
	if err != nil {
		return nil, err
	}

	sSlice := make([]*model.User, 0, len(msSlice))
	for _, ms := range msSlice {
		sSlice = append(sSlice, model.UserFromDomainModel(ms))
	}

	return sSlice, nil
}

// 追加
func (uu *userUsecase) FindUserById(ctx context.Context, id int) (*model.User, error) {
	ms, err := uu.svc.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.UserFromDomainModel(ms), nil
}
