package service

import (
	"app/db"
	"app/model"
	"app/repository"

	"math/rand"
	"time"
	"fmt"

	"github.com/gocraft/dbr"
)

type (
	IMongon interface {
		GetMongon() (*model.Mongon, error)
	}

	Mongon struct {
		repository repository.IMongon
		session *dbr.Session
	}
)

func NewMongon() IMongon {
	return &Mongon{
		repository: repository.NewMongon(),
		session: db.GetSession(),
	}
}

func (s *Mongon) GetMongon() (*model.Mongon, error) {
	l := &model.Mongon{}
	_, err := s.session.Select("id").From("Mongons").Load(l)

	idsLength := &model.Count{}
	d, err := s.session.Select("count(*) as value").From("Mongons").Load(idsLength)

	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", d)
	}

	rand.Seed(time.Now().UnixNano())
	randId := rand.Int63n(idsLength.Value) + 1

	m, err := s.repository.MongonByRandom(randId)
	if err != nil {
		return nil, err
	}

	return m, nil
}
