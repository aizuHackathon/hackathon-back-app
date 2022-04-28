package model

type User struct {
	ID     int64
	Name   string
	Height float32
	Weight float32
	Sex    int64
	Old    int64
}

type Users []User
