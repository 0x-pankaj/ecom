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

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	useHandler := user.NewHandler()
	useHandler.RegisterRoutes(subrouter)

	log.Println("server runing on port: 9090")

	return http.ListenAndServe(s.addr, router)
}
