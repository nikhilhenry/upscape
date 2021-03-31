package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"upscape/controllers"
	"upscape/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the application")

	// initilise database

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// connect to DB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err.Error())
	}
	// set database
	clientDatabase := client.Database(os.Getenv("DATABASE"))
	// when disconnected

	fmt.Println("Connected to Atlas successfully")
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// instantiate router
	router := gin.Default()

	// declare routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Upscape!"})
	})
	router.POST("/api/user", controllers.CreateUser(clientDatabase))
	router.POST("/api/login", controllers.LoginUser(clientDatabase))
	router.PUT("/api/user", middlewares.IsAuthenticated, controllers.UpdateUser(clientDatabase))

	// @todo remove this request
	router.GET("/howdy", middlewares.IsAuthenticated, func(c *gin.Context) {
		user_id := c.MustGet("user_id").(string)
		c.JSON(http.StatusOK, gin.H{"msg": "Hello!", "user_id": user_id})
	})

	router.Run()
}
