package config

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnapramodaradhi/expense-tracker-api/internal/handler"
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
	v1.HandleFunc("/auth/signup", util.Make(authHandler.Signup))

	log.Printf("Listening on port %v\n", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, router))
}
