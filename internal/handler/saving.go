package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/krishnapramodaradhi/expense-tracker-api/internal/entity"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/util"
)

type SavingHandler struct {
	db *sql.DB
}

func NewSavingHandler(db *sql.DB) *SavingHandler {
	return &SavingHandler{db: db}
}

func (h *SavingHandler) CreateSavings(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value("userId")
	var req []entity.SavingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	if err := h.addSavings(req, userId.(string)); err != nil {
		return err
	}
	if err := h.updateBudget(userId.(string)); err != nil {
		return err
	}
	return util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Savings added successfully"})
}

func (h *SavingHandler) updateBudget(userId string) error {
	_, err := h.db.Exec(util.UPDATE_BUDGET, userId)
	return err
}

func (h *SavingHandler) addSavings(savings []entity.SavingRequest, userId string) error {
	return util.BulkInsert(h.db, "savings", func(stmt *sql.Stmt) error {
		for _, saving := range savings {
			if _, err := stmt.Exec(saving.Title, saving.SType, saving.Amount, userId, time.Now()); err != nil {
				return err
			}
		}
		return nil
	}, "title", "type", "amount", "user_id", "created_at")
}
