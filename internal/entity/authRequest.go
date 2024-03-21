package entity

type AuthRequest struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	MonthlySalary int    `json:"monthlySalary"`
}
