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
		GetIDByNamePass(name, pass string) (*int64, error)
		CountName(name string) (int64, error)
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
		_, err := r.session.Select("*").From("users").Where("id IN ?", ids).Load(m)
		if err != nil {
			return nil, fmt.Errorf("fetch error :%v", err)
		}
		return m, nil
	}
	_, err := r.session.Select("*").From("users").Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}

func (r *User) GetIDByNamePass(name, pass string) (*int64, error) {
	user := &model.User{}
	if len(name) == 0 {
		return nil, fmt.Errorf("name is required")
	}
	if len(pass) == 0 {
		return nil, fmt.Errorf("pass is required")
	}
	ret, err := r.session.Select("*").From("users").
		Where("name = ?", name).
		Where("pass = ?", pass).
		Load(user)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	if ret == 0 {
		return nil, fmt.Errorf("couldn't find user by name: %v, pass: %v", name, pass)
	}

	return &(*user).ID, nil
}

func (r *User) CountName(name string) (int64, error) {
	var count = int64(1)
	_, err := r.session.Select("COUNT(*)").From("users").
		Where("name = ?", name).
		Load(&count)
	if err != nil {
		return 1, fmt.Errorf("fetch error :%v", err)
	}
	return count, nil
}

func (r *User) Create(m *model.CreateUser) error {
	_, err := r.session.InsertInto("users").
		Columns("name", "height", "sex", "old", "pass").
		Record(m).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
