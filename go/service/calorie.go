package service

import (
	"app/model"
	"app/repository"
)

type (
	ICalorie interface {
		GetCaloriesByUserIdCalorieType(id int64, calorieType model.CalorieType) (*model.Calories, error)
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

func (s *Calorie) GetCaloriesByUserIdCalorieType(id int64, calorieType model.CalorieType) (*model.Calories, error) {
	m, err := s.repository.ByUserIdCalorieType(id, calorieType)
	if err != nil {
		return nil, err
	}

	return m, nil
}
