package service

import (
	"app/model"
	"app/repository"
)

type (
	IUser interface {
		GetAllMongons() (*model.Mongons, error)
	}

	User struct {
		repository repository.IMongon
	}
)

func NewMongon() IMongon {
	return &Mongon{
		repository: repository.NewMongon(),
	}
}

func (s *Mongon) GetAllMongons() (*model.Mongons, error) {
	m, err := s.repository.ByIDs([]int64{})
	if err != nil {
		return nil, err
	}

	return m, nil
}