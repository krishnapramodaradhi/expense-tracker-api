package entity

type SavingRequest struct {
	Title  string `json:"title"`
	SType  string `json:"type"`
	Amount int    `json:"amount"`
}
