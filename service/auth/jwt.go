package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/arhantbararia/ecom_api/utils"
)

func CreateJWT(secret []byte, userId string) (string, error) {
	
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

func GetUserIDFromToken(r *http.Request) (string, error) {
	
	secret := []byte(utils.GetEnv("JWT_SECRET", "temp_secret"))
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	userId := claims["userId"].(string)
	return userId, nil

}
