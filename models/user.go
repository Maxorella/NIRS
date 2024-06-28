package models

import "time"

type Users struct {
	UserID    int64
	UserName  string
	Email     string
	DateBirth time.Time
	RegDate   time.Time
}
