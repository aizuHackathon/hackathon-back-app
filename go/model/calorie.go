package model

import "time"

type Calorie struct {
	ID           int64
	User_ID      int64
	Calorie_type int64
	Value        int64
	Created_at   time.Time
}

type Calories []Calorie
