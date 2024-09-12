package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/arhantbararia/ecom_api/utils"
)

func CreateJWT(secret []byte, userId string) (string, error) {
	//Create a new JWT token
	//Return the token and error if any
	expiration := utils.GetEnvInt("JWT_EXPIRATION", 3600*24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"expiredAt": time.Now().Add(time.Duration(expiration)),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
