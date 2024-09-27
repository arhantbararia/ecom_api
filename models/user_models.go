package models

import "time"

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	CreatedAt   time.Time `json:"createdAt"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type RegisterUserPayload struct {
	FirstName string `json:"fistname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}


type LoginUserPayload struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}
