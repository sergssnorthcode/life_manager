package models

import "time"

type User struct {
	ID         int
	TelegramID int64
	Username   string
	FirstName  string
	SecondName string
	CreatedAt  time.Time
}
