package handler

import (
	"net/http"
	"net/http/httptest"
	"resto-service/api/model"
	"resto-service/internal/app/resto/handler/mocks"
	"testing"

	basehandler "github.com/correooke/MyMeal/common/handler"
	baserepo "github.com/correooke/MyMeal/common/repository"
	baserouter "github.com/correooke/MyMeal/common/router"
	baseservice "github.com/correooke/MyMeal/common/service"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// Test for the endpoint GET /restos
func TestGetAllRestos(t *testing.T) {

	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)
	mockCategories := []model.Resto{
		{ID: "1", Name: "Test Resto 1", Description: "This is a test resto 1"},
		{ID: "2", Name: "Test Resto 2", Description: "This is a test resto 2"},
	}
	var docs []interface{}
	for _, resto := range mockCategories {
		docs = append(docs, resto)
	}
	cursorMock, _ := mongo.NewCursorFromDocuments(docs, nil, nil)
	collMock.On("Find", mock.Anything).Return(cursorMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Resto](collMock)
	repo := baserepo.NewLoggingRepository[model.Resto](repoWithoutLog)

	service := baseservice.NewService[model.Resto](repo)
	handler := basehandler.NewCommonHandler[model.Resto](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Resto](r2, &handler, "restos")

	// 3. Create a GET request
	req := httptest.NewRequest(http.MethodGet, "/restos", nil)
	rr := httptest.NewRecorder()

	c := r2.NewContext(req, rr)
	handler.GetAll(c)

	// 5. Check the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Additional checks for the returned restos

	collMock.AssertNumberOfCalls(t, "Find", 1)

}
