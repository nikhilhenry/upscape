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
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate = validator.New()

// create user

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

// update user meta data

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

		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
		}

		// query database
		result := collection.FindOneAndUpdate(ctx,
			bson.M{},
			bson.M{"$set": userUpdate},
			&opt,
		)

		if err := result.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Fatal(err)
		}

		// decode result
		var updatedUser UserUpdate
		if err := result.Decode(&updatedUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, updatedUser)
	}
}

// update user password

type PasswordUpdate struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}

func UpdateUserPassword(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("users")

		var passwordUpdate PasswordUpdate

		// bind object
		if err := c.BindJSON(&passwordUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate input
		if validationErr := validate.Struct(passwordUpdate); validationErr != nil {
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

		// check if current password matches
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordUpdate.CurrentPassword))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Incorrect email or password."})
			return
		}

		// hash new password password
		bytes, err := bcrypt.GenerateFromPassword([]byte(passwordUpdate.NewPassword), 14)
		if err != nil {
			log.Fatal(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		hashedPassword := string(bytes)

		// save to database
		result, err := collection.UpdateOne(
			ctx,
			bson.M{},
			bson.M{"$set": bson.M{"password": hashedPassword}},
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Fatal(err)
		}

		if result.MatchedCount < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "user info doesnt exist"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})

	}
}

type UserMeta struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	ImgUrl string `json:"img_url,omitempty" bson:"img_url,omitempty"`
}

func GetUserMeta(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("users")

		// query Database
		cursor := collection.FindOne(ctx, bson.M{})

		var userMeta UserMeta
		cursor.Decode(&userMeta)

		c.JSON(http.StatusOK, userMeta)

	}
}
