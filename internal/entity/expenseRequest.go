package entity

type ExpenseRequest struct {
	Title       string `json:"title"`
	EType       string `json:"type"`
	PaymentMode string `json:"paymentMode"`
	PaymentTo   string `json:"paymentTo"`
	Amount      int    `json:"amount"`
	Reason      string `json:"reason"`
}
