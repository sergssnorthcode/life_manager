package models

import "time"

type User struct {
	ID         int
	TelegramID int64
	Username   string
	FirstName  string
	LastName   string
	CreatedAt  time.Time
}
