package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/0x-pankaj/ecom/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {

	//gorilla router to handle route
	router := mux.NewRouter()

	// creating subrouter for simplification
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// creating user handler
	userStore := user.NewStore(s.db)
	useHandler := user.NewHandler(userStore)
	useHandler.RegisterRoutes(subrouter)

	log.Println("server runing on port: 9090")

	return http.ListenAndServe(s.addr, router)
}
