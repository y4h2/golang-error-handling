package entity

import "time"

type Article struct {
	ID        string
	Title     string
	Content   string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
