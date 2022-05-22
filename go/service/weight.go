package service

import (
	"app/form"
	"app/model"
	"app/repository"
	"fmt"
)

type (
	IWeight interface {
		GetWeightsByUserID(id int64) (*model.Weights, error)
		Create(f *form.Weight, userId int64) error
	}

	Weight struct {
		repository repository.IWeight
	}
)

func NewWeight() IWeight {
	return &Weight{
		repository: repository.NewWeight(),
	}
}

func (s *Weight) GetWeightsByUserID(id int64) (*model.Weights, error) {
	m, err := s.repository.ByUserID(id)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Weight) Create(f *form.Weight, userId int64) error {
	m := &model.CreateWeight{
		User_ID: userId,
		Value:   f.Weight,
	}

	if err := s.repository.Create(m); err != nil {
		return fmt.Errorf("failed to create weight: %v, err: %w", m, err)
	}

	return nil
}
