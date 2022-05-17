package model

type User struct {
	ID     int64
	Name   string
	Height float32
	Sex    int64
	Old    int64
}

type Users []User

type CreateUser struct {
	Name   string
	Height float32
	Sex    int64
	Old    int64
	Pass   string
}
