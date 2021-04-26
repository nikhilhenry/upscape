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

// get tag
func GetTags(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tags")

		// query database
		// todo:add pagination and sort by ID for results
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var tags []models.Tag
		// loop through results
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var tag models.Tag
			cursor.Decode(&tag)
			tags = append(tags, tag)
		}

		if tags == nil {
			myslice := make([]string, 0)
			c.JSON(http.StatusOK, myslice)

			return
		}

		// send tags
		c.JSON(http.StatusOK, tags)
	}
}

// create tag
func CreateTag(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tags")

		// create new tag variable
		var tag models.Tag

		// bind tag
		if err := c.BindJSON(&tag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate input
		if validationErr := validate.Struct(tag); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// post to database
		result, insertErr := collection.InsertOne(ctx, tag)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		// get the document
		var createdTag models.Tag
		document := helpers.GetDocByID(collection, result.InsertedID.(primitive.ObjectID), &createdTag)

		// return the document
		c.JSON(http.StatusOK, document)
	}
}

func UpdateTag(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tags")

		// get tag ID
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		var tagUpdate models.Tag

		// bind oject
		if err := c.BindJSON(&tagUpdate); err != nil {
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
			bson.M{"$set": tagUpdate},
			&opts,
		)

		if err := result.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// decode result
		var updatedtag models.Tag
		if err := result.Decode(&updatedtag); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedtag)
	}
}

func DeleteTag(client *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// establish connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := client.Collection("tags")

		// get tag ID
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
