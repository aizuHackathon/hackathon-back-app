package response

import (
	"app/model"
	"time"
)

type (
	Weight struct {
		ID         int64     `json:"id"`
		User_ID    int64     `json:"user_id"`
		Value      float32   `json:"value"`
		Created_At time.Time `json:"create_at"`
	}

	Weights []Weight
)

func NewWeight(m *model.Weight) *Weight {
	if m == nil {
		return nil
	}

	return &Weight{
		ID:         m.ID,
		User_ID:    m.User_ID,
		Value:      m.Value,
		Created_At: m.Created_at,
	}
}

func NewWeights(m *model.Weights) *Weights {
	ret := Weights{}
	if m == nil {
		return nil
	}

	for _, v := range *m {
		r := NewWeight(&v)
		if r == nil {
			continue
		}
		ret = append(ret, *r)
	}

	return &ret
}
