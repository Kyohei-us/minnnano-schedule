package model

import (
	"minnnano-schedule/domain/model"
)

type UserID int

type UserSlice []*User

type User struct {
	ID   UserID `json:"id"`
	Name string `json:"name"`
}

func UserFromDomainModel(m *model.User) *User {
	s := &User{
		ID:   UserID(m.ID),
		Name: m.Name,
	}

	return s
}
