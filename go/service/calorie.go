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
		GetStatus(userId int64) (int64, error)
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

func (s *Calorie) GetStatus(userId int64) (int64, error) {
	mealCalorie, err := s.repository.AllCaloriesByUserIDAndType(userId, 0)
	if err != nil {
		return -1, err
	}
	exerciseCalorie, err := s.repository.AllCaloriesByUserIDAndType(userId, 1)
	if err != nil {
		return -1, err
	}
	// TODO: change border
	fmt.Println(mealCalorie.Value, exerciseCalorie.Value)
	if float32(exerciseCalorie.Value)*0.6+float32(mealCalorie.Value)*0.4 > 500 {
		return 0, nil
	} else {
		return 1, nil
	}
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
