package service

import (
	"app/model"
	"app/repository"
	"fmt"
	"math/rand"
	"time"
)

type (
	IKeihatu interface {
		GetKeihatu() (*model.Keihatu, error)
		IdByRandom() (int64, error)
	}
	Keihatu struct {
		repository repository.IKeihatu
	}
)

func NewKeihatu() IKeihatu {
	return &Keihatu{
		repository: repository.NewKeihatu(),
	}
}

func (s *Keihatu) GetKeihatu() (*model.Keihatu, error) {

	id, err := s.IdByRandom()
	if err != nil {

		return nil, fmt.Errorf("error IdByRandom: %v", err)
	}
	keihatu, err := s.repository.KeihatuByID(id)
	if err != nil {
		return nil, err
	}
	return keihatu, nil
}

//データベース上にあるIDをランダムで取得、配列を作成してその中から要素を一つだけ取得
func (s *Keihatu) IdByRandom() (int64, error) {
	m, err := s.repository.All()
	if err != nil {
		return -1, err
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(*m))
	keihatu := (*m)[index].ID

	return keihatu, nil
}
