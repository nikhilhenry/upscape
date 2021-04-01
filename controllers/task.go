package controllers

import (
	"context"
	"net/http"
	"time"
	"upscape/helpers"
	"upscape/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// get tasks
func GetTasks(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tasks")

		// query Database
		cursor, err := collection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var tasks []models.Task

		// loop through results
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var task models.Task
			cursor.Decode(&task)
			tasks = append(tasks, task)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// send result
		c.JSON(http.StatusOK, tasks)
	}
}

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

		// get document
		var createdTask models.Task
		document := helpers.GetDocByID(collection, result.InsertedID.(primitive.ObjectID), &createdTask)

		// return the document
		c.JSON(http.StatusOK, document)
	}
}
