package service

import (
	"app/model"
	"app/repository"
)

type (
	IWeight interface {
		GetWeightsByUserID(id int64) (*model.Weights, error)
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
