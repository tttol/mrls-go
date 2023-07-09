package model

import "time"

type MrDetail struct {
	Title     string
	WebUrl    string
	Author    User
	Upvotes   int
	Labels    []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
