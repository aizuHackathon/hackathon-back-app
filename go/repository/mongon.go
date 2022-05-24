package repository

import (
	"app/db"
	"app/model"
	"fmt"
	"math/rand"
	"time"

	"github.com/gocraft/dbr"
)

type (
	IMongon interface {
		MongonByRandom(id int64) (*model.Mongon, error)
		RandomId() (int64, error)
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

func (r *Mongon) MongonByRandom(id int64) (*model.Mongon, error) {
	m := &model.Mongon{}
	_, err := r.session.Select("*").From("Mongons").Where("id = ?", id).Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	
	return m, nil
}

func (r *Mongon) RandomId() (int64, error) {
	idsLength := &model.Count{}
	_, err := r.session.Select("count(*) as value").From("Mongons").Load(idsLength)

	if err != nil {
		return 0, fmt.Errorf("fetch error :%v", err)
	}

	rand.Seed(time.Now().UnixNano())
	randId := rand.Int63n(idsLength.Value) + 1

	return randId, nil
}
