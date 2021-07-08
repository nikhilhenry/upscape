package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"upscape/helpers"
	"upscape/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// get objective
func GetObjectives(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("objectives")

		// initialize find options
		findOptions := options.Find()

		// get sort key and field
		if sort := c.Query("sort"); sort != "" {
			
			if field := c.Query("field"); field == "created_at" {
				if sort == "asc"{
					findOptions.SetSort(bson.D{{"created_at",1}})
				}else{
					findOptions.SetSort(bson.D{{"created_at",-1}})
				}
			}else if field == "scheduled_for"{
				if sort == "asc"{
					findOptions.SetSort(bson.D{{"scheduled_for",1}})
				}else{
					findOptions.SetSort(bson.D{{"scheduled_for",-1}})
				}
			}
		}

		// get page and limit
		page,_ := strconv.Atoi(c.Query("page"))
		limit,_ := strconv.Atoi(c.Query("limit"))

		if page==0 || limit ==0 {
			page = 1
			limit = 0
		}
		findOptions.SetSkip(int64((page-1)*limit))
		findOptions.SetLimit(int64(limit))

		// query database
		// todo:add pagination and sort by ID for results
		cursor, err := collection.Find(ctx, bson.M{},findOptions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var objectives []models.Objective
		// loop through results
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var objective models.Objective
			cursor.Decode(&objective)
			objectives = append(objectives, objective)
		}

		if objectives == nil {
			myslice := make([]string, 0)
			c.JSON(http.StatusOK, myslice)

			return
		}

		// send objectives
		c.JSON(http.StatusOK, objectives)
	}
}

// create objective
func CreateObjective(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("objectives")

		// create new objective variable
		var objective models.Objective

		// bind objective
		if err := c.BindJSON(&objective); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate input
		if validationErr := validate.Struct(objective); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// assign timestamp
		objective.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		// post to database
		result, insertErr := collection.InsertOne(ctx, objective)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		// get the document
		var createdObjective models.Objective
		document := helpers.GetDocByID(collection, result.InsertedID.(primitive.ObjectID), &createdObjective)

		// return the document
		c.JSON(http.StatusOK, document)
	}
}

func UpdateObjective(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("objectives")

		// get objective ID
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		var objectiveUpdate models.Objective

		// bind oject
		if err := c.BindJSON(&objectiveUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		after := options.After
		opts := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
		}

		// query database
		result := collection.FindOneAndUpdate(
			ctx,
			bson.M{"_id": id},
			bson.M{"$set": objectiveUpdate},
			&opts,
		)

		if err := result.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// decode result
		var updatedobjective models.Objective
		if err := result.Decode(&updatedobjective); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedobjective)
	}
}

func DeleteObjective(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("objectives")

		// get objective ID
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
