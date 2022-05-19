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
		GetNameCount(f *form.UserName) (int64, error)
		GetUserIdByNamePass(name, pass string) (*int64, error)
		Create(f *form.User) error
	}

	User struct {
		userRepository   repository.IUser
		weightRepository repository.IWeight
	}
)

func NewUser() IUser {
	return &User{
		userRepository:   repository.NewUser(),
		weightRepository: repository.NewWeight(),
	}
}

func (s *User) GetAllUsers() (*model.Users, error) {
	m, err := s.userRepository.ByIDs([]int64{})
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *User) GetNameCount(f *form.UserName) (int64, error) {
	count, err := s.userRepository.CountName(f.Name)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (s *User) GetUserIdByNamePass(name, pass string) (*int64, error) {
	id, err := s.userRepository.GetIDByNamePass(name, pass)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id name: %v, pass: %v, err: %w", name, pass, err)
	}

	return id, nil
}

func (s *User) Create(f *form.User) error {
	name := f.Name
	pass := f.Pass
	m := &model.CreateUser{
		Name:   f.Name,
		Height: f.Height,
		Sex:    f.Sex,
		Old:    f.Old,
		Pass:   f.Pass,
	}

	if err := s.userRepository.Create(m); err != nil {
		return fmt.Errorf("failed to create user: %v, err: %w", m, err)
	}

	id, err := s.userRepository.GetIDByNamePass(name, pass)
	if err != nil {
		return fmt.Errorf("failed to get userId, name: %v, pass: %v, err: %w", name, pass, err)
	}

	createWeight := &model.CreateWeight{
		User_ID: *id,
		Value:   f.Weight,
	}

	if err := s.weightRepository.Create(createWeight); err != nil {
		return fmt.Errorf("failed to create weight: %v, err: %w", createWeight, err)
	}

	return nil
}
