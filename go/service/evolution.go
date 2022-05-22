package service

import (
	"app/model"
	"app/repository"
)

type (
	IEvolution interface {
		GetStatus(userId int64) (int64, error)
	}

	Evolution struct {
		repository repository.IEvolution
	}
)

func NewEvolution() IEvolution {
	return &Evolution{
		repository: repository.NewEvolution(),
	}
}

func checkEvolutionStatus(workoutCalorie, mealCalorie *model.Status) int64 {
	if float32(workoutCalorie.Value)*0.7+float32(mealCalorie.Value)*0.3 > 11340 {
		return 1
	}
	return 0
}

func (s *Evolution) GetStatus(userId int64) (int64, error) {
	mealCalorie, err := s.repository.AllCaloriesByUserIdCalorieType(userId, model.CalorieTypeMeal)
	if err != nil {
		return -1, err
	}
	workoutCalorie, err := s.repository.AllCaloriesByUserIdCalorieType(userId, model.CalorieTypeWorkout)
	if err != nil {
		return -1, err
	}
	return checkEvolutionStatus(workoutCalorie, mealCalorie), nil
}
