package handler

import (
	"category-service/api/model"
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

// Test for the endpoint GET /categories/{id}
func TestGetCategoryById(t *testing.T) {

	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)
	mockCategory := model.Category{ID: "1", Name: "Test Category", Description: "This is a test category"}

	singleResultMock := mongo.NewSingleResultFromDocument(mockCategory, nil, nil)
	collMock.On("FindOne", mock.Anything).Return(singleResultMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Category](collMock)
	repo := baserepo.NewLoggingRepository[model.Category](repoWithoutLog)

	service := baseservice.NewService[model.Category](repo)
	handler := basehandler.NewCommonHandler[model.Category](service)

	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.Category](r, &handler, "categories")

	// 3. Create a GET request
	req := httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	rec := httptest.NewRecorder()

	c := r.NewContext(req, rec)
	c.SetPath("/categories/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// 4. Call the corresponding handler
	if err := handler.GetByID(c); err != nil {
		t.Fatalf("Handler returned an error: %v", err)
	}

	// 5. Check the response
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Additional checks for the returned category

	collMock.AssertNumberOfCalls(t, "FindOne", 1)
}
