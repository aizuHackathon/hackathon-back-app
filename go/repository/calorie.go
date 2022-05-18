package repository

import (
	"app/db"
	"app/model"
	"fmt"

	"github.com/gocraft/dbr"
)

type (
	ICalorie interface {
		ByUserID(id int64, calorieType int64) (*model.Calories, error)
	}

	Calorie struct {
		session *dbr.Session
	}
)

func NewCalorie() ICalorie {
	return &Calorie{
		session: db.GetSession(),
	}
}

func (r *Calorie) ByUserID(id int64, calorieType int64) (*model.Calories, error) {
	m := &model.Calories{}
	_, err := r.session.Select("*").From("calories").
		Where("user_id = ?", id).
		Where("calorie_type = ?", calorieType).
		OrderAsc("created_at").
		Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}
