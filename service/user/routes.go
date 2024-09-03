package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

//each service will be of type Handler
type Handler struct {

}




func NewHandler() *Handler {
	return &Handler{}
}


func (h *Handler ) RegisterRoutes (router *mux.Router ) {
	router.HandleFunc("/login" , h.LoginHandle ).Methods("POST")
	router.HandleFunc("/register" , h.RegisterHandle).Methods("POST")
}

func (h *Handler ) LoginHandle (w http.ResponseWriter , r *http.Request ){
	
}

func (h *Handler) RegisterHandle (w http.ResponseWriter , r *http.Request ) {
	
}