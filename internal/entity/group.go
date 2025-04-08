package entity

import "time"

type Group struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	DueDay    int    `json:"due_day"`
	CreatedAt time.Time `json:"created_at"`
	Users     []User `json:"users"`
}
