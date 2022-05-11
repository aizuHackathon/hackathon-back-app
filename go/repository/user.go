package repository

import (
	"app/db"
	"app/model"
	"fmt"

	"github.com/gocraft/dbr"
)

type (
	IUser interface {
		ByIDs(ids []int64) (*model.Users, error)
		Create(m *model.CreateUser) error
	}

	User struct {
		session *dbr.Session
	}
)

func NewUser() IUser {
	return &User{
		session: db.GetSession(),
	}
}

func (r *User) ByIDs(ids []int64) (*model.Users, error) {
	m := &model.Users{}
	if len(ids) != 0 {
		_, err := r.session.Select("*").From("Users").Where("id IN ?", ids).Load(m)
		if err != nil {
			return nil, fmt.Errorf("fetch error :%v", err)
		}
		return m, nil
	}
	_, err := r.session.Select("*").From("Users").Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}

func (r *User) Create(m *model.CreateUser) error {
	_, err := r.session.InsertInto("users").
		Columns("name", "height", "weight", "sex", "old").
		Record(m).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
