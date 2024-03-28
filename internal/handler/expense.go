package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/krishnapramodaradhi/expense-tracker-api/internal/entity"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/util"
)

type ExpenseHandler struct {
	db *sql.DB
}

func NewExpenseHandler(db *sql.DB) *ExpenseHandler {
	return &ExpenseHandler{db: db}
}

func (h *ExpenseHandler) CreateExpenses(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value("userId")
	var req []entity.ExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	if err := util.BulkInsert(h.db, "expenses", func(stmt *sql.Stmt) error {
		for _, expense := range req {
			if _, err := stmt.Exec(expense.Title, expense.EType, expense.PaymentMode, expense.PaymentTo, expense.Amount, expense.Reason, userId, time.Now()); err != nil {
				return err
			}
		}
		return nil
	}, "title", "type", "payment_mode", "payment_to", "amount", "reason", "user_id", "created_at"); err != nil {
		return err
	}
	return util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Expenses added successfully"})
}
