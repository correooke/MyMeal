package handler

import (
	"dish-service/api/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"dish-service/internal/app/dish/handler/mocks"

	basehandler "github.com/correooke/MyMeal/common/handler"
	baserepo "github.com/correooke/MyMeal/common/repository"
	baserouter "github.com/correooke/MyMeal/common/router"
	baseservice "github.com/correooke/MyMeal/common/service"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// Test for the endpoint GET /dishes/{id}
func TestGetDishById(t *testing.T) {

	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)
	mockDish := model.Dish{ID: "1", Name: "Test Dish", Description: "This is a test dish"}

	singleResultMock := mongo.NewSingleResultFromDocument(mockDish, nil, nil)
	collMock.On("FindOne", mock.Anything).Return(singleResultMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Dish](collMock)
	repo := baserepo.NewLoggingRepository[model.Dish](repoWithoutLog)

	service := baseservice.NewService[model.Dish](repo)
	handler := basehandler.NewCommonHandler[model.Dish](service)

	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.Dish](r, &handler, "dishes")

	// 3. Create a GET request
	req := httptest.NewRequest(http.MethodGet, "/dishes/1", nil)
	rec := httptest.NewRecorder()

	c := r.NewContext(req, rec)
	c.SetPath("/dishes/:id")
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

	// TODO: Additional checks for the returned dish

	collMock.AssertNumberOfCalls(t, "FindOne", 1)
}
