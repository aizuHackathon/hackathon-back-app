package repository

import (
	"app/db"
	"app/model"
	"fmt"

	"github.com/gocraft/dbr"
)

type (
	IKeihatu interface {
		KeihatuByIDs(ids []int64) (*model.Keihatus, error)
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

func (r *Keihatu) KeihatuByIDs(ids []int64) (*model.Keihatus, error) {
	m := &model.Keihatus{}
	if len(ids) != 0 {
		_, err := r.session.Select("*").From("Keihatus").Where("id IN ?", ids).Load(m)
		if err != nil {
			return nil, fmt.Errorf("fetch error :%v", err)
		}
		return m, nil
	}
	_, err := r.session.Select("*").From("Keihatus").Load(m)
	if err != nil {
		return nil, fmt.Errorf("fetch error :%v", err)
	}
	return m, nil
}
