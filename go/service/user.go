package service

import (
	"app/model"
	"app/repository"
)

type (
	IUser interface {
		GetAllUsers() (*model.Users, error)
	}

	User struct {
		repository repository.IUser
	}
)

func NewUser() IUser {
	return &User{
		repository: repository.NewUser(),
	}
}

func (s *User) GetAllUsers() (*model.Users, error) {
	m, err := s.repository.ByIDs([]int64{})
	if err != nil {
		return nil, err
	}

	return m, nil
}
