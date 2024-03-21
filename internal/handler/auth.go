package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/krishnapramodaradhi/expense-tracker-api/internal/entity"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/util"
)

type AuthHandler struct{}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) error {
	var authRequest entity.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		log.Println("An error occured while decoding request body", err)
		return err
	}
	return util.WriteJSON(w, http.StatusOK, authRequest)
}
