package model

import "time"

type Timer struct {
	Id        string
	Title     string
	DueDate   time.Time
	CreatedAt time.Time
}
