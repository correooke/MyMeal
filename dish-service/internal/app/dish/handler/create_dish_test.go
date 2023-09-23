package handler

import (
	"bytes"
	"dish-service/api/model"
	"dish-service/internal/app/dish/handler/mocks"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	basehandler "github.com/correooke/MyMeal/common/handler"
	baserepo "github.com/correooke/MyMeal/common/repository"
	baserouter "github.com/correooke/MyMeal/common/router"
	baseservice "github.com/correooke/MyMeal/common/service"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// Test para el endpoint POST localhost:8080/dishes
func TestCreateDish(t *testing.T) {
	// 1. Configurar la base de datos en memoria (puedes usar un paquete como go-memdb o similar).
	collMock := new(mocks.CollectionMock)
	mockDish := model.Dish{ID: "1", Name: "Test Dish", Description: "This is a test dish"}

	singleResultMock := mongo.NewSingleResultFromDocument(mockDish, nil, nil)

	collMock.On("InsertOne", mock.Anything).Return(&mongo.InsertOneResult{InsertedID: "1"}, nil)
	collMock.On("FindOne", mock.Anything).Return(singleResultMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Dish](collMock)
	repo := baserepo.NewLoggingRepository[model.Dish](repoWithoutLog)

	service := baseservice.NewService[model.Dish](repo)
	handler := basehandler.NewCommonHandler[model.Dish](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Dish](r2, &handler, "dishes")

	// 3. Crear una solicitud POST con una categoría de ejemplo.
	dish := model.Dish{
		ID:          "1",
		Name:        "Test Dish",
		Description: "This is a test dish",
		// ... otros campos ...
	}
	body, _ := json.Marshal(dish)
	req := httptest.NewRequest(http.MethodGet, "/dishes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	c := r2.NewContext(req, rr)

	handler.Create(c)

	// 5. Comprobar la respuesta.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 6. Comprobar que la categoría se ha insertado en la base de datos en memoria.
	insertedDish, _ := repo.GetByID(context.TODO(), dish.ID)
	if insertedDish.Name != dish.Name {
		t.Errorf("Inserted dish does not match: got %v want %v", insertedDish.Name, dish.Name)
	}

	collMock.AssertNumberOfCalls(t, "InsertOne", 1)
	collMock.AssertCalled(t, "FindOne", mock.Anything)
}
