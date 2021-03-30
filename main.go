package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the application")

	// instantiate router
	router := gin.Default()

	// declare routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Upscape!"})
	})

	router.Run()
}
