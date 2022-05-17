package repository

import (
	"app/db"
	"app/model"
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/labstack/gommon/log"
)

type (
	IWeight interface {
		ByUserID(id int64) (*model.Weights, error)
		Create(m *model.CreateWeight) error
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
	_, err := r.session.Select("*").From("weights").Where("user_id = ?", id).Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	log.Infof("get value: %v", m)
	return m, nil
}

func (r *Weight) Create(m *model.CreateWeight) error {
	_, err := r.session.InsertInto("weights").
		Columns("user_id", "value").
		Record(m).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
