package service

import (
	"app/form"
	"app/model"
	"app/repository"
	"fmt"
)

type (
	ICalorie interface {
		GetCaloriesByUserID(id int64, calorieType int64) (*model.Calories, error)
		Create(f *form.Calorie, userId int64, calorieType int64) error
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

func (s *Calorie) Create(f *form.Calorie, userId int64, calorieType int64) error {
	m := &model.CreateCalorie{
		User_ID:      userId,
		Calorie_type: calorieType,
		Value:        f.Calorie,
	}

	if err := s.repository.Create(m); err != nil {
		return fmt.Errorf("failed to create calorie: %v, err: %w", m, err)
	}

	return nil
}