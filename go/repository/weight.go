package repository

import (
	"app/db"
	"app/model"
	"fmt"

	"github.com/gocraft/dbr"
)

type (
	IWeight interface {
		ByUserID(id int64) (*model.Weights, error)
	}

	Weight struct {
		session *dbr.Session
	}
)

func NewWeight() IWeight {
	return &Weight{
		session: db.GetSession(),
	}
}

func (r *Weight) ByUserID(id int64) (*model.Weights, error) {
	m := &model.Weights{}
	_, err := r.session.Select("*").From("Weights").Where("user_id = ?", id).Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}
