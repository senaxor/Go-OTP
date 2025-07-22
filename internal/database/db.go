package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "OTP/internal/models"
)

var DB *mongo.Database

func InitDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    DB = client.Database("otp_project")

    models.InitUserCollection(DB)
}
