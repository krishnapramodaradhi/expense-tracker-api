package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/krishnapramodaradhi/expense-tracker-api/internal/entity"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) error {
	var req entity.AuthRequest
	var userId string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPassword)
	if err := h.db.QueryRow(util.SIGNUP, req.FirstName, req.LastName, req.Email, req.Password, req.MonthlySalary, time.Now()).Scan(&userId); err != nil {
		return err
	}
	token, err := util.GenerateToken(userId)
	if err != nil {
		return err
	}
	return util.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var req entity.AuthRequest
	var userId string
	var password string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	if err := h.db.QueryRow(util.FETCH_USER_BY_ID, req.Email).Scan(&userId, &password); err != nil {
		return fmt.Errorf("Auth Failed")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return fmt.Errorf("Auth Failed")
	}
	token, err := util.GenerateToken(userId)
	if err != nil {
		return err
	}
	return util.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
