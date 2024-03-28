package entity

import "time"

type Salary struct {
	Id        string    `json:"id"`
	Amount    int       `json:"amount"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
