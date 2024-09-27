package user

import "github.com/gorilla/mux"

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.LoginHandle).Methods("POST")
	router.HandleFunc("/register", h.RegisterHandle).Methods("POST")
	router.HandleFunc("/user", h.GetUserData).Methods("GET")
}
