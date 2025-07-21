package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Phone     string             `bson:"phone" json:"phone"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

var userCollection *mongo.Collection

func InitUserCollection(db *mongo.Database) {
	userCollection = db.Collection("users")
}

func FindOrCreateUser(phone string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user already exists
	var user User
	err := userCollection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user)
	if err == nil {
		return &user, nil // found
	}
	if err != mongo.ErrNoDocuments {
		return nil, err // DB error
	}

	// Create new user
	newUser := User{
		Phone:     phone,
		CreatedAt: time.Now(),
	}

	res, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	newUser.ID = res.InsertedID.(primitive.ObjectID)
	return &newUser, nil
}

func GetUserByID(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var user User
	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
