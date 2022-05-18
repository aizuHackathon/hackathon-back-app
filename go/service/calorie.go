package service

import (
	"app/model"
	"app/repository"
)

type (
	ICalorie interface {
		GetCaloriesByUserID(id int64, calorieType int64) (*model.Calories, error)
	}

	Calorie struct {
		repository repository.ICalorie
	}
)

func NewCalorie() ICalorie {
	return &Calorie{
		repository: repository.NewCalorie(),
	}
}

func (s *Calorie) GetCaloriesByUserID(id int64, calorieType int64) (*model.Calories, error) {
	m, err := s.repository.ByUserID(id, calorieType)
	if err != nil {
		return nil, err
	}

	return m, nil
}
