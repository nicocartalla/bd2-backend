package models

import (
	"time"
)

type Notification struct {
	ID          int
	UserId 		int
	MatchId 	int
	NotificationTime time.Time
	NotificationMethod string
}
type Data struct {
	Email    string    `json:"email"`
	Subject  string `json:"subject"`
	Body string `json:"body"`
}


type NotifData struct {
	Data Data
	NotifModel Notification
}
