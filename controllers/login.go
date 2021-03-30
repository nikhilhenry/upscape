package controllers

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	Password string `json:"password" bson:"password"`
}

func generateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)

	// define claims
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// sign token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return tokenString, nil
}
