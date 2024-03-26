package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/krishnapramodaradhi/expense-tracker-api/internal/entity"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/util"
)

type AuthHandler struct{ db *sql.DB }

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) error {
	var authRequest entity.AuthRequest
	var userId string
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		log.Println("An error occured while decoding request body", err)
		return err
	}
	if err := h.db.QueryRow("insert into users (first_name, last_name, email, password, updated_at) values ($1, $2, $3, $4, $5) returning id", authRequest.FirstName, authRequest.LastName, authRequest.Email, authRequest.Password, time.Now()).Scan(&userId); err != nil {
		return err
	}
	return util.WriteJSON(w, http.StatusOK, userId)
}
