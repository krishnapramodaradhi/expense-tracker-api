package config

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/handler"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/middleware"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/util"
)

type Server struct {
	port string
	db   *sql.DB
}

func NewServer(port string, db *sql.DB) *Server {
	return &Server{port: port, db: db}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()

	authHandler := handler.NewAuthHandler(s.db)
	savingHandler := handler.NewSavingHandler(s.db)
	expenseHandler := handler.NewExpenseHandler(s.db)

	v1.HandleFunc("/auth/signup", util.Make(authHandler.Signup))

	v1Protected := v1.PathPrefix("").Subrouter()
	v1Protected.Use(middleware.TokenValidation)
	v1Protected.HandleFunc("/savings", util.Make(savingHandler.CreateSavings)).Methods(http.MethodPost)

	v1Protected.HandleFunc("/expenses", util.Make(expenseHandler.CreateExpenses)).Methods(http.MethodPost)

	log.Printf("Listening on port %v\n", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, router))
}
