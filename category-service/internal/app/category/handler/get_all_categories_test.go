package handler

import (
	"category-service/api/model"
	"category-service/internal/app/category/handler/mocks"
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

// Test for the endpoint GET /categories
func TestGetAllCategories(t *testing.T) {

	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)
	mockCategories := []model.Category{
		{ID: "1", Name: "Test Category 1", Description: "This is a test category 1"},
		{ID: "2", Name: "Test Category 2", Description: "This is a test category 2"},
	}
	var docs []interface{}
	for _, category := range mockCategories {
		docs = append(docs, category)
	}
	cursorMock, _ := mongo.NewCursorFromDocuments(docs, nil, nil)
	collMock.On("Find", mock.Anything).Return(cursorMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Category](collMock)
	repo := baserepo.NewLoggingRepository[model.Category](repoWithoutLog)

	service := baseservice.NewService[model.Category](repo)
	handler := basehandler.NewCommonHandler[model.Category](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Category](r2, &handler, "categories")

	// 3. Create a GET request
	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	rr := httptest.NewRecorder()

	c := r2.NewContext(req, rr)
	handler.GetAll(c)

	// 5. Check the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Additional checks for the returned categories

	collMock.AssertNumberOfCalls(t, "Find", 1)

}
