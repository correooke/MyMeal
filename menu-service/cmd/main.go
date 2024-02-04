package main

import (
	"context"
	"fmt"
	"log"
	"menu-service/api/model"
	"menu-service/api/router"
	"net/http"

	"github.com/correooke/MyMeal/common/db"
	basehandler "github.com/correooke/MyMeal/common/handler"
	baserepo "github.com/correooke/MyMeal/common/repository"
	baseservice "github.com/correooke/MyMeal/common/service"
)

func main() {
	cli, err := db.ConnectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	mongoCollection := cli.Database("mymealdb").Collection("menues")
	collection := db.NewRealCollection(mongoCollection)
	repo := baserepo.NewMongoRepository[model.Menu](collection)
	srv := baseservice.NewService[model.Menu](repo)
	handler := basehandler.NewCommonHandler[model.Menu](srv)

	r := router.NewRouter(&handler)

	defer cli.Disconnect(context.TODO())

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
