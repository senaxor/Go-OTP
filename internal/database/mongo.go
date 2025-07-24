package database

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "OTP/internal/models"
)

var DB *mongo.Database

func InitMongo() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017" // fallback for local dev
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    DB = client.Database("otp_project")

    models.InitUserCollection(DB)
}
