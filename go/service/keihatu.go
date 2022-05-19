package service

import (
	"app/model"
	"app/repository"
)

type (
	IKeihatu interface {
		GetKeihatuByID() (*model.Keihatu, error)
	}
	Keihatu struct {
		repository repository.IKeihatu
	}
)

func NewKeihatu() IKeihatu {
	return &Keihatu{
		repository: repository.NewKeihatu(),
	}
}

func (s *Keihatu) GetKeihatuByID() (*model.Keihatu, error) {
	m, err := s.repository.KeihatuByID()
	if err != nil {
		return nil, err
	}
	return m, nil
}
