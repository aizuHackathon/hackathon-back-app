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
	IKeihatu interface {
		KeihatuByID() (*model.Keihatu, error)
	}

	Keihatu struct {
		session *dbr.Session
	}
)

func NewKeihatu() IKeihatu {
	return &Keihatu{
		session: db.GetSession(),
	}
}

func (r *Keihatu) KeihatuByID() (*model.Keihatu, error) {
	m := &model.Keihatu{}
	idsLength := &model.Count{}
	_, err := r.session.Select("count(*) as value").From("Keihatu").Load(idsLength)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	rand.Seed(time.Now().UnixNano())
	randId := rand.Int63n(idsLength.Value)
	_, error := r.session.Select("*").From("Keihatu").Where("id =?", randId).Load(m)
	if error != nil {
		return nil, fmt.Errorf("fetch error :%v", error)
	}
	return m, nil
}
