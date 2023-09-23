package handler

import (
	"bytes"
	"resto-service/api/model"
	"resto-service/internal/app/resto/handler/mocks"
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

// Test para el endpoint POST localhost:8080/restos
func TestCreateResto(t *testing.T) {
	// 1. Configurar la base de datos en memoria (puedes usar un paquete como go-memdb o similar).
	collMock := new(mocks.CollectionMock)
	mockResto := model.Resto{ID: "1", Name: "Test Resto", Description: "This is a test resto"}

	singleResultMock := mongo.NewSingleResultFromDocument(mockResto, nil, nil)

	collMock.On("InsertOne", mock.Anything).Return(&mongo.InsertOneResult{InsertedID: "1"}, nil)
	collMock.On("FindOne", mock.Anything).Return(singleResultMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Resto](collMock)
	repo := baserepo.NewLoggingRepository[model.Resto](repoWithoutLog)

	service := baseservice.NewService[model.Resto](repo)
	handler := basehandler.NewCommonHandler[model.Resto](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Resto](r2, &handler, "restos")

	// 3. Crear una solicitud POST con una categoría de ejemplo.
	resto := model.Resto{
		ID:          "1",
		Name:        "Test Resto",
		Description: "This is a test resto",
		// ... otros campos ...
	}
	body, _ := json.Marshal(resto)
	req := httptest.NewRequest(http.MethodGet, "/restos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	c := r2.NewContext(req, rr)

	handler.Create(c)

	// 5. Comprobar la respuesta.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 6. Comprobar que la categoría se ha insertado en la base de datos en memoria.
	insertedResto, _ := repo.GetByID(context.TODO(), resto.ID)
	if insertedResto.Name != resto.Name {
		t.Errorf("Inserted resto does not match: got %v want %v", insertedResto.Name, resto.Name)
	}

	collMock.AssertNumberOfCalls(t, "InsertOne", 1)
	collMock.AssertCalled(t, "FindOne", mock.Anything)
}
