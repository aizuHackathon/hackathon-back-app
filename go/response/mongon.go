package response

import "app/model"

type (
	Mongon struct {
		Id     int64   `json:"id"`
		Mongon   string  `json:"mongon"`
	}

	Mongons []Mongon
)

func NewMongon(m *model.Mongon) *Mongon {
	if m == nil {
		return nil
	}

	return &User{
		Id:     m.Id,
		Mongon:   m.Mongon,
	}
}

func NewMongons(m *model.Mongons) *Mongons {
	ret := Mongons{}
	if m == nil {
		return nil
	}

	for _, v := range *m {
		r := NewMongon(&v)
		if r == nil {
			continue
		}
		ret = append(ret, *r)
	}

	return &ret
}
