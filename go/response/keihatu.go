package response

import (
	"app/model"
)

type (
	Keihatu struct {
		ID    int64  `json:"id"`
		Value string `json:"value"`
	}
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
