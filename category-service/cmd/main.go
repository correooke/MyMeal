package main

import (
	"category-service/api/model"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/correooke/MyMeal/common/db"
	basehandler "github.com/correooke/MyMeal/common/handler"
	baserepo "github.com/correooke/MyMeal/common/repository"
	baserouter "github.com/correooke/MyMeal/common/router"
	baseservice "github.com/correooke/MyMeal/common/service"
)

func main() {
	/*fmt.Println("Initiating Categories API...")
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
	*/
	client2, err := db.ConnectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	repo2 := baserepo.NewMongoRepository[model.Category](client2, "categoriesdb", "categories")
	service2 := baseservice.NewService[model.Category](repo2)
	categoryHandler2 := basehandler.NewCommonHandler[model.Category](service2)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Category](r2, &categoryHandler2, "categories")
	baserouter.AddIsAlive(r2)

	defer client2.Disconnect(context.TODO())

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r2)
}
