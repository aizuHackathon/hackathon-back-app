package model

import "time"

type Weight struct {
	ID         int64
	User_ID    int64
	Value      float32
	Created_at time.Time
}

type Weights []Weight

type CreateWeight struct {
	User_ID int64
	Value   float32
}
