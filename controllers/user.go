package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"upscape/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate = validator.New()

func CreateUser(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("users")

		// check if user document already exists
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(ctx)
		if cursor.RemainingBatchLength() > 0 {
			fmt.Println("user exists")
			c.JSON(http.StatusConflict, gin.H{"error": "user exists"})
			return
		}

		// create a user variable
		var user models.User

		// bind object
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate input
		if validationErr := validate.Struct(user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// hash password
		bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			log.Fatal(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.Password = string(bytes)

		// post do Database
		result, insertErr := collection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		// return the object
		c.JSON(http.StatusOK, result)
	}
}

type UserUpdate struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	ImgUrl string `json:"img_url,omitempty" bson:"img_url,omitempty"`
}

func UpdateUser(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("users")

		// create user-update variable
		var userUpdate UserUpdate

		// bind object
		if err := c.BindJSON(&userUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate reqeust
		if validationErr := validate.Struct(userUpdate); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		// query database
		result, err := collection.UpdateOne(
			ctx,
			bson.M{},
			bson.M{"$set": userUpdate},
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Fatal(err)
		}

		fmt.Println(result.ModifiedCount)

		c.JSON(http.StatusOK, result)
	}
}
