package model

import "time"

type CalorieType int

const (
	CalorieTypeMeal CalorieType = iota
	CalorieTypeWorkout
)

type Calorie struct {
	ID           int64
	User_ID      int64
	Calorie_type CalorieType
	Value        int64
	Created_at   time.Time
}

type Calories []Calorie

type CreateCalorie struct {
	User_ID      int64
	Calorie_type CalorieType
	Value        int64
}

func (c CalorieType) IsInvalid() bool {
	switch c {
	case CalorieTypeMeal, CalorieTypeWorkout:
		return false
	}
	return true
}
