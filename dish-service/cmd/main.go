package main

import (
	"context"
	"dish-service/api/router"
	"dish-service/internal/app/dish/handler"
	"dish-service/internal/app/dish/repository"
	"dish-service/internal/app/dish/service"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Setup MongoDB connection
	fmt.Println("Initiating Dishes API...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	repo := repository.NewMongoDishRepository(client, "dishesdb")
	logRepo := repository.NewLogDishRepositoryDecorator(repo)
	service := service.NewDishService(logRepo)
	dishHandler := handler.NewDishHandler(service)

	// Setup routes
	r := router.NewRouter(dishHandler)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
