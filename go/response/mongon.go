package response

import "app/model"

type (
	Mongon struct {
		ID     int64   `json:"id"`
		Mongon   string  `json:"mongon"`
	}
)

func NewMongon(m *model.Mongon) *Mongon {
	if m == nil {
		return nil
	}

	return &Mongon{
		ID:     m.ID,
		Mongon: m.Mongon,
	}
}
