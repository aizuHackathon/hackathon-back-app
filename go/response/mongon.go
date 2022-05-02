package response

import "app/model"

type (
	Mongon struct {
		ID     int64   `json:"id"`
		Mongon   string  `json:"mongon"`
	}

	Mongons []Mongon
)

func NewMongon(m *model.Mongon) *Mongon {
	if m == nil {
		return nil
	}

	return &Mongon{
		ID:     m.ID,
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
