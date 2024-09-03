package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arhantbararia/ecom_api/service/user"
)

type APIServer struct {
	address string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		address: addr,
	}

}

func (s *APIServer) Run() error {

	//setting up router
	router := mux.NewRouter()

	//setting up subrouter for modularity
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//starting user service
	log.Println("starting user service ")
	userService := user.NewHandler()
	userService.RegisterRoutes(subrouter)

	log.Printf("\n\nE-Comm API Server Running\n")
	log.Println("Serving on", s.address)

	return http.ListenAndServe(s.address, subrouter)

}
