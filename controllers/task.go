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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// get tasks
func GetTasks(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tasks")

		// @todo: get tasks based on query params
		// var tasks []models.Task

		// get param from query
		dateRange := c.Query("range")

		// sort all results by id
		opts := options.Find().SetSort(bson.M{"id": 1})

		// if date range is today
		if dateRange == "today" {
			// query Database for uncompleted tasks
			cursor, err := collection.Find(ctx, bson.M{"completed": false}, opts)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			tasks := getDocsFromCursor(cursor)
			// send result
			c.JSON(http.StatusOK, tasks)
		}
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

func getDocsFromCursor(cursor *mongo.Cursor) []models.Task {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tasks []models.Task

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	return tasks
}
