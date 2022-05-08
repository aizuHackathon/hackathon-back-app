package response

import (
	"app/model"
)

type (
	Keihatu struct {
		ID    int64  `json:"id"`
		Value string `json:"value"`
	}
	Keihatus []Keihatu
)

func NewKeihatu(m *model.Keihatu) *Keihatu {
	if m == nil {
		return nil
	}

	return &Keihatu{
		ID:    m.ID,
		Value: m.Value,
	}
}

func NewKeihatus(m *model.Keihatus) *Keihatus {
	ret := Keihatus{}
	if m == nil {
		return nil
	}

	for _, v := range *m {
		r := NewKeihatu(&v)
		if r == nil {
			continue
		}
		ret = append(ret, *r)
	}

	return &ret
}
