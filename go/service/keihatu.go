package service

import (
	"app/model"
	"app/repository"
)

type (
	IKeihatu interface {
		GetAllKeihatus() (*model.Keihatus, error)
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

func (s *Keihatu) GetAllKeihatus() (*model.Keihatus, error) {
	m, err := s.repository.KeihatuByIDs([]int64{})
	if err != nil {
		return nil, err
	}
	return m, nil
}
