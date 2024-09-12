package user

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/arhantbararia/ecom_api/models"
	"github.com/arhantbararia/ecom_api/service/auth"
	"github.com/google/uuid"
)

func checkUserExists(db *sql.DB, email string) (string, error) {
	//query to get user and return its user_id if exists
	query := "SELECT id FROM users WHERE email = ?"
	var id string
	err := db.QueryRow(query, email).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error getting user: %v", err)
	}
	return id, nil


	
}

func CreateNewUserTable(db *sql.DB) error {
	log.Printf("Creating User Table")
	query := `CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36) PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating user table: %v", err)
	}
	log.Println("User Table Created!")
	return nil
}

func createNewUser(db *sql.DB, payload models.RegisterUserPayload) error {
	// now create a user_id using uuid

	uuid, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error creating new user id: %v", err)
	}
	//convert uuid to string
	userID := uuid.String()

	//hash Password
	hashedPassword, err := auth.HashedPassword(payload.Password)
	if err != nil {
		log.Println("Error Hashing the password.")
	}

	//create new user
	user := models.User{
		ID:        userID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
	//insert user into db
	query := "INSERT INTO users (id, first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("error inserting new user: %v", err)
	}
	return nil

}


func GetUser(db *sql.DB , userID string , user *models.User ) error {
	//query to get user and return its user_id if exists
	query := "SELECT * FROM users WHERE id = ?"
	
	err := db.QueryRow(query, userID).Scan(
		&user.ID , 
		&user.FirstName , 
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}
	return nil

}


