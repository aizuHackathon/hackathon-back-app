package repository

import (
	"app/db"
	"app/model"
	"fmt"
	"math/rand"

	"github.com/gocraft/dbr"
)

type (
	IMongon interface {
		ByIDs(ids []int64) (*model.Mongons, error)
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

func (r *Mongon) ByIDs(ids []int64) (*model.Mongons, error) {
	m := &model.Mongons{}
	if len(ids) != 0 {
		_, err := r.session.Select("*").From("Mongons").Where("id IN ?", ids).Load(m)
		if err != nil {
			return nil, fmt.Errorf("fetch error :%v", err)
		}
		return m, nil
	}
	_, err := r.session.Select("*").From("Mongons").Where("id = ?",rand.Intn(2) + 1).Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	
	return m, nil
}
