package controllers

import (
	"context"
	"fmt"
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

		// get param from query
		dateRange := c.Query("range")

		if len(dateRange) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "please specify date range"})
			return
		}

		// sort all results by id
		opts := options.Find().SetSort(bson.M{"id": 1})

		// @todo: query for completed_at instead of created_at


		// Set location
		location,_ := time.LoadLocation("Asia/Kolkata")

		if dateRange == "tomorrow" {
			// define filter
			lowTime := time.Now()
			filter := bson.D{{"created_at", bson.D{{"$gt", lowTime}}}}

			// query Database for uncompleted tasks
			cursor, err := collection.Find(ctx, filter, opts)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			tasks := getDocsFromCursor(cursor)
			// send result
			c.JSON(http.StatusOK, tasks)
			return
		}

		// if date range is today
		if dateRange == "today" {

			// define filter
			t := time.Now()
			highTime := time.Date(t.Year(), t.Month(), t.Day(), 24, 59, 59, 100000000, location)
			lowTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, location)
			filter := bson.D{{"created_at", bson.D{{"$gt", lowTime.UTC()}, {"$lt", highTime.UTC()}}}}

			// query Database for uncompleted tasks
			cursor, err := collection.Find(ctx, filter, opts)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			tasks := getDocsFromCursor(cursor)
			// send result
			c.JSON(http.StatusOK, tasks)
			return
		}

		// for custom date
		fmt.Println(dateRange)
		lowTime, _ := time.ParseInLocation("01/02/06", dateRange, location)
		highTime := lowTime.Add(time.Hour*23 + time.Minute*59 + time.Second*59)
		fmt.Println(lowTime)
		filter := bson.D{{"created_at", bson.D{{"$gt", lowTime.UTC()}, {"$lt", highTime.UTC()}}}}
		cursor, err := collection.Find(ctx, filter, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks := getDocsFromCursor(cursor)
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


		// Set location
		location,_ := time.LoadLocation("Asia/Kolkata")
		t := time.Now().In(location)

		// assign timestamp
		if task.IsTomorrow {
			task.CreatedAt = time.Date(t.Year(),t.Month(),t.Day()+1,0,0,0,0,location).UTC()
		} else {
			task.CreatedAt, _ = time.Parse(time.RFC3339, t.Format(time.RFC3339))
		}

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

func UpdateTask(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tasks")

		// get task ID
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		var taskUpdate models.TaskUpdate

		// bind oject
		if err := c.BindJSON(&taskUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if taskUpdate.Completed != nil {
			taskUpdate.CompletedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		}

		after := options.After
		opts := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
		}

		// query database
		result := collection.FindOneAndUpdate(
			ctx,
			bson.M{"_id": id},
			bson.M{"$set": taskUpdate},
			&opts,
		)

		if err := result.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// decode result
		var updatedTask models.Task
		if err := result.Decode(&updatedTask); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedTask)
	}
}

func DeleteTask(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tasks")

		// get task ID
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		// query Database
		result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if result.DeletedCount > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
		}

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

	if tasks == nil {
		myslice := make([]models.Task, 0)

		return myslice
	}

	return tasks
}
