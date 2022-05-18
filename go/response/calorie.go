package response

import (
	"app/model"
	"time"
)

type (
	Calorie struct {
		Value      int64     `json:"value"`
		Created_At time.Time `json:"create_at"`
	}

	Calories []Calorie

	Status struct {
		Value string `json:"value"`
	}
)

func NewCalorie(m *model.Calorie) *Calorie {
	if m == nil {
		return nil
	}

	return &Calorie{
		Value:      m.Value,
		Created_At: m.Created_at,
	}
}

func NewCalories(m *model.Calories) *Calories {
	ret := Calories{}
	if m == nil {
		return nil
	}

	for _, v := range *m {
		r := NewCalorie(&v)
		if r == nil {
			continue
		}
		ret = append(ret, *r)
	}

	return &ret
}

func NewStatus(status int64) *Status {
	if status == 0 {
		return &Status{
			Value: "first",
		}
	} else if status == 1 {
		return &Status{
			Value: "second",
		}
	} else {
		return &Status{
			Value: "third",
		}
	}
}
