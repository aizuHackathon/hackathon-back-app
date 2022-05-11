package service

import (
	"app/form"
	"app/model"
	"app/repository"
	"fmt"
)

type (
	IUser interface {
		GetAllUsers() (*model.Users, error)
		Create(f *form.User) error
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

func (s *User) Create(f *form.User) error {
	m := &model.CreateUser{
		// ID:     f.ID,
		Name:   f.Name,
		Height: f.Height,
		Weight: f.Weight,
		Sex:    f.Sex,
		Old:    f.Old,
	}

	if err := s.repository.Create(m); err != nil {
		return fmt.Errorf("failed to create user: %v, err: %w", m, err)
	}

	return nil
}
