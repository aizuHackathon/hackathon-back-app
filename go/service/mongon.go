package service

import (
	"app/model"
	"app/repository"
)

type (
	IMongon interface {
		GetMongon() (*model.Mongons, error)
	}

	Mongon struct {
		repository repository.IMongon
	}
)

func NewMongon() IMongon {
	return &Mongon{
		repository: repository.NewMongon(),
	}
}

func (s *Mongon) GetMongon() (*model.Mongons, error) {
	m, err := s.repository.ByIDs([]int64{})
	if err != nil {
		return nil, err
	}

	return m, nil
}
