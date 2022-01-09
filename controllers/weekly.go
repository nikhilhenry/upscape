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

func GetWeeklyItems(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		//	establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("actionables")

		//	create filter and sort
		opts := options.Find().SetSort(bson.M{"id": 1})
		filter := bson.M{"inbox":false}

		//	query database
		cursor, err := collection.Find(ctx, filter, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		weeklyItems := getActionablesFromCursor(cursor)
		c.JSON(http.StatusOK, weeklyItems)
	}
}

func CreateWeeklyItem(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("actionables")

		//	create new actionable variable
		var weeklyItem models.Actionable

		if err := c.BindJSON(&weeklyItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//	set to weekly
		weeklyItem.Inbox = false

		//	validate input
		if validationErr := validate.Struct(weeklyItem); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}
		//	post to database
		result, insertErr := collection.InsertOne(ctx, weeklyItem)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		}

		//	get document
		var createdWeeklyItem models.Actionable
		document := helpers.GetDocByID(collection, result.InsertedID.(primitive.ObjectID), &createdWeeklyItem)

		//	return the document
		c.JSON(http.StatusOK, document)
	}
}

func UpdateActionableItem(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) { // establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("actionables")

		// get actionable ID
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		var actionableUpdate models.ActionableUpdate

		// bind object
		if err := c.BindJSON(&actionableUpdate); err != nil {
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
			bson.M{"$set": actionableUpdate},
			&opts,
		)

		if err := result.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// decode result
		var updatedActionable models.Actionable
		if err := result.Decode(&updatedActionable); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedActionable)
	}
}
