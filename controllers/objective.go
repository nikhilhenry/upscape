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

// get objective
func GetObjectives(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("objectives")

		// query database
		// todo:add pagination and sort by ID for results
		cursor, err := collection.Find(ctx, bson.M{})
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
