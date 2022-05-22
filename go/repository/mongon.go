package repository

import (
	"app/db"
	"app/model"
	"fmt"

	"github.com/gocraft/dbr"
)

type (
	IMongon interface {
		MongonByRandom(ids int64) (*model.Mongon, error)
	}

	Mongon struct {
		session *dbr.Session
	}
)

func NewMongon() IMongon {
	return &Mongon{
		session: db.GetSession(),
	}
}

func (r *Mongon) MongonByRandom(ids int64) (*model.Mongon, error) {
	m := &model.Mongon{}
	_, err := r.session.Select("*").From("Mongons").Where("id = ?", ids).Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	
	return m, nil
}
