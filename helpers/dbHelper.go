package helpers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @todo add error handling

func GetDocByID(collection *mongo.Collection, ID primitive.ObjectID, document interface{}) interface{} {

	// estbalish connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// query for document
	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(document)
	if err != nil {

		return err
	}
	return document
}

func GetDocFromCursor(cursor *mongo.Cursor, document interface{}) []interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var documents []interface{}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		cursor.Decode(document)
		documents = append(documents, document)
	}

	return documents
}
