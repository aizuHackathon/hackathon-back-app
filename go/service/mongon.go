package service

import (
	"app/model"
	"app/repository"
)

type (
	IMongon interface {
		GetMongon() (*model.Mongon, error)
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

func (s *Mongon) GetMongon() (*model.Mongon, error) {
	Id, err := s.repository.RandomId()
	if err != nil {
		return nil, err
	}

	m, err := s.repository.MongonByRandom(Id)
	if err != nil {
		return nil, err
	}

	return m, nil
}
