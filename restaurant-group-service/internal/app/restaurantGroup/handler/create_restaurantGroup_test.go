package handler

import (
	"bytes"
	"restaurantGroup-service/api/model"
	"restaurantGroup-service/internal/app/restaurantGroup/handler/mocks"
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

// Test para el endpoint POST localhost:8080/restaurantGroups
func TestCreateRestaurantGroup(t *testing.T) {
	// 1. Configurar la base de datos en memoria (puedes usar un paquete como go-memdb o similar).
	collMock := new(mocks.CollectionMock)
	mockRestaurantGroup := model.RestaurantGroup{ID: "1", Name: "Test RestaurantGroup", Description: "This is a test restaurantGroup"}

	singleResultMock := mongo.NewSingleResultFromDocument(mockRestaurantGroup, nil, nil)

	collMock.On("InsertOne", mock.Anything).Return(&mongo.InsertOneResult{InsertedID: "1"}, nil)
	collMock.On("FindOne", mock.Anything).Return(singleResultMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.RestaurantGroup](collMock)
	repo := baserepo.NewLoggingRepository[model.RestaurantGroup](repoWithoutLog)

	service := baseservice.NewService[model.RestaurantGroup](repo)
	handler := basehandler.NewCommonHandler[model.RestaurantGroup](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.RestaurantGroup](r2, &handler, "restaurantGroups")

	// 3. Crear una solicitud POST con una categoría de ejemplo.
	restaurantGroup := model.RestaurantGroup{
		ID:          "1",
		Name:        "Test RestaurantGroup",
		Description: "This is a test restaurantGroup",
		// ... otros campos ...
	}
	body, _ := json.Marshal(restaurantGroup)
	req := httptest.NewRequest(http.MethodGet, "/restaurantGroups", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	c := r2.NewContext(req, rr)

	handler.Create(c)

	// 5. Comprobar la respuesta.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 6. Comprobar que la categoría se ha insertado en la base de datos en memoria.
	insertedRestaurantGroup, _ := repo.GetByID(context.TODO(), restaurantGroup.ID)
	if insertedRestaurantGroup.Name != restaurantGroup.Name {
		t.Errorf("Inserted restaurantGroup does not match: got %v want %v", insertedRestaurantGroup.Name, restaurantGroup.Name)
	}

	collMock.AssertNumberOfCalls(t, "InsertOne", 1)
	collMock.AssertCalled(t, "FindOne", mock.Anything)
}
