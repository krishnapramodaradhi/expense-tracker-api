package entity

import "time"

type Expense struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	EType       string    `json:"type"`
	PaymentMode string    `json:"paymentMode"`
	PaymentTo   string    `json:"paymentTo"`
	Amount      int       `json:"amount"`
	Reason      string    `json:"reason"`
	UserId      string    `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
