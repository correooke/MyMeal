package main

import (
	"category-service/api/router"
	"category-service/internal/app/category/handler"
	"category-service/internal/app/category/repository"
	"category-service/internal/app/category/service"
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Initiating Categories API...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	repo := repository.NewMongoCategoryRepository(client, "categoriesdb")
	service := service.NewCategoryService(repo)
	categoryHandler := handler.NewCategoryHandler(service)

	r := router.NewRouter(categoryHandler)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
