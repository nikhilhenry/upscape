package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
	"upscape/helpers"
	"upscape/models"
)

func GetInboxItems(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		//	establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("actionables")

		// create filter and sort
		opts := options.Find().SetSort(bson.M{"id": 1})
		filter := bson.D{{"inbox", true}}

		//	query Database
		cursor, err := collection.Find(ctx, filter, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		inboxItems := getActionablesFromCursor(cursor)
		c.JSON(http.StatusOK, inboxItems)
		return
	}
}

func CreateInboxItem(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		//	establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("actionables")

		//	create new task variable
		var inboxItem models.Actionable

		if err := c.BindJSON(&inboxItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// set to inbox
		inboxItem.Inbox = true

		//	validate input
		if validationErr := validate.Struct(inboxItem); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		//	post to database
		result, insertErr := collection.InsertOne(ctx, inboxItem)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		//	get document
		var createdInboxItem models.Actionable
		document := helpers.GetDocByID(collection, result.InsertedID.(primitive.ObjectID), &createdInboxItem)

		//	return the document
		c.JSON(http.StatusOK, document)
	}
}

func DeleteInboxItem(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		//	establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("actionables")

		//	get item id
		_id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		//	query database
		result, err := collection.DeleteOne(ctx, bson.M{"_id": _id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if result.DeletedCount > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
		}
	}
}

func getActionablesFromCursor(cursor *mongo.Cursor) []models.Actionable {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var actionables []models.Actionable

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var actionable models.Actionable
		cursor.Decode(&actionable)
		actionables = append(actionables, actionable)
	}

	if actionables == nil {
		myslice := make([]models.Actionable, 0)

		return myslice
	}

	return actionables
}
