package entity

import "time"

type Saving struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Amount    int       `json:"amount"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
