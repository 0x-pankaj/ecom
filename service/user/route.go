package user

import (
	"fmt"
	"net/http"

	"github.com/0x-pankaj/ecom/types"
	"github.com/0x-pankaj/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (r *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", r.handleLogin).Methods("POST")
	router.HandleFunc("/register", r.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//get user and check if already registered
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user with email already exist"))
	}
	//hash password

	//create user
	h.store.CreateUserWithEmail(types.User{
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  payload.Password,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

}
