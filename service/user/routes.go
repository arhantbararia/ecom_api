package user

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/arhantbararia/ecom_api/models"
	"github.com/arhantbararia/ecom_api/service/auth"
	"github.com/arhantbararia/ecom_api/utils"
	"github.com/gorilla/mux"
)

// each service will be of type Handler
type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {

	//Creating User Table
	err := CreateNewUserTable(db)
	if err != nil {
		log.Fatalf("Error creating user table: %v", err)
	}

	return &Handler{
		db: db,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.LoginHandle).Methods("POST")
	router.HandleFunc("/register", h.RegisterHandle).Methods("POST")
}

func (h *Handler) LoginHandle(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no Request body found"))
		return
	}

	var payload models.LoginUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)

	}

	// check if user already exists
	userID, err := checkUserExists(h.db, payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if userID == "" {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("no user found for this email-id"))
		return
	}

	//get user from userid
	var user models.User
	err = GetUser(h.db, userID, &user)
	if err != nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("some Error Occurred: %v", err))
		return
	}

	// check if password is correct
	if !auth.ComparePassword(user.Password, payload.Password) {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("wrong password"))
		return
	}

	secret := utils.GetEnv("JWT_SECRET", "temp_secret")

	JWT_TOKEN, err := auth.CreateJWT([]byte(secret), userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("some error Occurred"))
	}

	utils.WriteJSON(w, http.StatusAccepted, map[string]string{"map": JWT_TOKEN})
	

}

func (h *Handler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no request body found"))
		return
	}

	var payload models.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if user already exists
	userId, err := checkUserExists(h.db, payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if userId != "" {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user already exists"))
		return
	}

	// create new user
	if err := createNewUser(h.db, payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "User Created!")

}
