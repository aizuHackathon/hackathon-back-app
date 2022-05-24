package repository

import (
	"app/db"
	"app/model"
	"fmt"

	"github.com/gocraft/dbr"
)

type (
	IKeihatu interface {
		KeihatuByID(id int64) (*model.Keihatu, error)
		All() (*model.Keihatus, error)
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

//データベース上にあるIDを取得するためのメソッド
func (r *Keihatu) All() (*model.Keihatus, error) {
	m := &model.Keihatus{}
	_, err := r.session.Select("*").From("keihatu").Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}

//選択されたIDの啓発を取得するメソッド
func (r *Keihatu) KeihatuByID(id int64) (*model.Keihatu, error) {
	m := &model.Keihatu{}
	_, error := r.session.Select("*").From("Keihatu").Where("id = ?", id).Load(m)
	if error != nil {
		return nil, fmt.Errorf("fetch error :%v", error)
	}
	return m, nil
}
