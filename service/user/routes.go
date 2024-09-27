package user

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/arhantbararia/ecom_api/models"
	"github.com/arhantbararia/ecom_api/service/auth"
	"github.com/arhantbararia/ecom_api/utils"
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

func (h *Handler) GetUserData(w http.ResponseWriter, r *http.Request) {
	//we are using JWT
	//client will send JWT token in header
	//we will extract the token and get the user_id from it
	//then we will get the user from the user_id

	userID, err := auth.GetUserIDFromToken(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	// get user from db
	var user models.User
	err = GetUser(h.db, userID, &user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("some error occurred"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
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
	if !auth.ComparePassword(payload.Password, user.Password) {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("wrong password"))
		return
	}

	secret := utils.GetEnv("JWT_SECRET", "temp_secret")

	JWT_TOKEN, err := auth.CreateJWT([]byte(secret), userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("some error Occurred"))
	}

	utils.WriteJSON(w, http.StatusAccepted, map[string]string{"Token": JWT_TOKEN})

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

	// check if user already exists by the email
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

func (h *Handler) UpdateUserData(w http.ResponseWriter, r *http.Request) {

	userID, err := auth.GetUserIDFromToken(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	var updated_user models.UpdateUserPayload
	if err := utils.ParseJson(r, &updated_user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)

	}

	err = UpdateUser(h.db, userID, updated_user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusAccepted, "User Updated Successfully")

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromToken(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	err = DeleteUser(h.db, userID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusAccepted, "User Deleted Successfully")
}


func (h* Handler) UpdatePassword ( w http.ResponseWriter , r *http.Request ) {

}
