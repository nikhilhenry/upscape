package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"upscape/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Password string `json:"password" bson:"password"`
}

func generateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// define claims
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	fmt.Println(os.Getenv("SECRET"))

	// sign token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return tokenString, nil
}

func LoginUser(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {

		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("users")

		var loginRequest LoginRequest

		// bind object
		if err := c.BindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate request
		validate := validator.New()
		if validationErr := validate.Struct(loginRequest); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// query database
		var user models.User
		err := collection.FindOne(ctx, bson.M{}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login failed"})
			return
		}

		// check password hash
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Incorrect email or password."})
			return
		}

		// generate token
		token, err := generateToken(user.ID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
