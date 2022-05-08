package response

import "app/model"

type (
	Weight struct {
		ID      int64 `json:"id"`
		User_ID int64 `json:"user_id"`
		Value   float32 `json:"value"`
	}

	Weights []Weight
)

func NewWeight(m *model.Weight) *Weight {
	if m == nil {
		return nil
	}

	return &Weight{
		ID:      m.ID,
		User_ID: m.User_ID,
		Value:   m.Value,
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
