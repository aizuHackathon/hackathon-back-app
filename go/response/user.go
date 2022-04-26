package response

import "app/model"

type (
	User struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	Users []User
)

func NewUser(m *model.User) *User {
	if m == nil {
		return nil
	}

	return &User{
		ID:   m.ID,
		Name: m.Name,
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
