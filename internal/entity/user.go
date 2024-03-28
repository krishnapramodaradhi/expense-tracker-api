package entity

import "time"

type User struct {
	Id             string    `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	MonthlySalary  int       `json:"monthlySalary"`
	BudgetPerMonth int       `json:"budgetPerMonth"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
