package handler

import (
	"dish-service/api/model"
	"dish-service/internal/app/dish/handler/mocks"
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

// Test for the endpoint GET /dishes
func TestGetAllCategories(t *testing.T) {

	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)
	mockCategories := []model.Dish{
		{ID: "1", Name: "Test Dish 1", Description: "This is a test dish 1"},
		{ID: "2", Name: "Test Dish 2", Description: "This is a test dish 2"},
	}
	var docs []interface{}
	for _, dish := range mockCategories {
		docs = append(docs, dish)
	}
	cursorMock, _ := mongo.NewCursorFromDocuments(docs, nil, nil)
	collMock.On("Find", mock.Anything).Return(cursorMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Dish](collMock)
	repo := baserepo.NewLoggingRepository[model.Dish](repoWithoutLog)

	service := baseservice.NewService[model.Dish](repo)
	handler := basehandler.NewCommonHandler[model.Dish](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Dish](r2, &handler, "dishes")

	// 3. Create a GET request
	req := httptest.NewRequest(http.MethodGet, "/dishes", nil)
	rr := httptest.NewRecorder()

	c := r2.NewContext(req, rr)
	handler.GetAll(c)

	// 5. Check the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Additional checks for the returned dishes

	collMock.AssertNumberOfCalls(t, "Find", 1)

}
