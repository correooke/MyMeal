package handler

import (
	"bytes"
	"category-service/api/model"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"category-service/internal/app/category/handler/mocks"

	basehandler "github.com/correooke/MyMeal/common/handler"
	baserepo "github.com/correooke/MyMeal/common/repository"
	baserouter "github.com/correooke/MyMeal/common/router"
	baseservice "github.com/correooke/MyMeal/common/service"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// Test para el endpoint POST localhost:8080/categories
func TestCreateCategory(t *testing.T) {
	// 1. Configurar la base de datos en memoria (puedes usar un paquete como go-memdb o similar).
	collMock := new(mocks.CollectionMock)
	mockCategory := model.Category{ID: "1", Name: "Test Category", Description: "This is a test category"}

	singleResultMock := mongo.NewSingleResultFromDocument(mockCategory, nil, nil)

	collMock.On("InsertOne", mock.Anything).Return(&mongo.InsertOneResult{InsertedID: "1"}, nil)
	collMock.On("FindOne", mock.Anything).Return(singleResultMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Category](collMock)
	repo := baserepo.NewLoggingRepository[model.Category](repoWithoutLog)

	service := baseservice.NewService[model.Category](repo)
	handler := basehandler.NewCommonHandler[model.Category](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Category](r2, &handler, "categories")

	// 3. Crear una solicitud POST con una categoría de ejemplo.
	category := model.Category{
		ID:          "1",
		Name:        "Test Category",
		Description: "This is a test category",
		// ... otros campos ...
	}
	body, _ := json.Marshal(category)
	req, _ := http.NewRequest("POST", "/categories", bytes.NewBuffer(body))

	// 4. Llamar al handler correspondiente.
	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	// 5. Comprobar la respuesta.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 6. Comprobar que la categoría se ha insertado en la base de datos en memoria.
	insertedCategory, _ := repo.GetByID(context.TODO(), category.ID)
	if insertedCategory.Name != category.Name {
		t.Errorf("Inserted category does not match: got %v want %v", insertedCategory.Name, category.Name)
	}

	collMock.AssertNumberOfCalls(t, "InsertOne", 1)
	collMock.AssertCalled(t, "FindOne", mock.Anything)
}
