package repository

import (
	"app/db"
	"app/model"
	"fmt"
	"github.com/gocraft/dbr"
)

type (
	IEvolution interface {
		AllCaloriesByUserIdCalorieType(id int64, calorieType model.CalorieType) (*model.Status, error)
	}

	Evolution struct {
		session *dbr.Session
	}
)

func NewEvolution() IEvolution {
	return &Evolution{
		session: db.GetSession(),
	}
}

func (r *Evolution) AllCaloriesByUserIdCalorieType(id int64, calorieType model.CalorieType) (*model.Status, error) {
	m := &model.Status{}
	t, err := r.session.Select("sum(value) as value").From("calories").
		Where("user_id = ?", id).
		Where("calorie_type = ?", calorieType).
		Rows()
	for t.Next() {
		err = t.Scan(&m.Value)
		if err != nil {
			return nil, fmt.Errorf("fetch error :%v", err)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}
