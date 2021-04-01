package controllers

import (
	"context"
	"net/http"
	"time"
	"upscape/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// create task
func CreateTask(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tasks")

		// create new task variable
		var task models.Task

		// bind object
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate input
		if validationErr := validate.Struct(task); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// assign timestamp
		task.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		// post to database
		result, insertErr := collection.InsertOne(ctx, task)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		// return the objectID
		c.JSON(http.StatusOK, result)
	}
}
