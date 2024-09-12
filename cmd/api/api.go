package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arhantbararia/ecom_api/middleware"
	"github.com/arhantbararia/ecom_api/service/user"
)

type APIServer struct {
	address string

	db *sql.DB
}

func NewAPIServer(addr string  , db *sql.DB) *APIServer {
	return &APIServer{
		address: addr,
		db: db,
	}

}

func (s *APIServer) Run() error {

	//setting up router
	router := mux.NewRouter()

	//setting up subrouter for modularity
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//setting up middleware
	middleware_stack := middleware.CreateStack(
		middleware.LogRequest,
	)

	//starting user service
	log.Println("starting user service ")
	userService := user.NewHandler(s.db)
	userService.RegisterRoutes(subrouter)

	log.Printf("\n\nE-Comm API Server Running\n")
	log.Println("Serving on", s.address)

	server := http.Server{
		Addr:    s.address,
		Handler: middleware_stack(subrouter),
	}

	return server.ListenAndServe()
	

}
