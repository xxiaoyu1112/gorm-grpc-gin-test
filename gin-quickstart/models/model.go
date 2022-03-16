package models

import "time"

type Todo struct {
	ID   int       `json:"id"`
	Item int       `json:"item"`
	Time time.Time `json:"time"`
}
