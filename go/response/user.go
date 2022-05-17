package response

import "app/model"

type (
	User struct {
		ID     int64   `json:"id"`
		Name   string  `json:"name"`
		Height float32 `json:"height"`
		Sex    int64   `json:"sex"`
		Old    int64   `json:"old"`
	}

	Users []User
)

func NewUser(m *model.User) *User {
	if m == nil {
		return nil
	}

	return &User{
		ID:     m.ID,
		Name:   m.Name,
		Height: m.Height,
		Sex:    m.Sex,
		Old:    m.Old,
	}
}

func NewUsers(m *model.Users) *Users {
	ret := Users{}
	if m == nil {
		return nil
	}

	for _, v := range *m {
		r := NewUser(&v)
		if r == nil {
			continue
		}
		ret = append(ret, *r)
	}

	return &ret
}
