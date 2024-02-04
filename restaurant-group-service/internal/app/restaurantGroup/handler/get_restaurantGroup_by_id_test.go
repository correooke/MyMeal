package handler

import (
	"restaurantGroup-service/api/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"restaurantGroup-service/internal/app/restaurantGroup/handler/mocks"

	basehandler "github.com/correooke/MyMeal/common/handler"
	baserepo "github.com/correooke/MyMeal/common/repository"
	baserouter "github.com/correooke/MyMeal/common/router"
	baseservice "github.com/correooke/MyMeal/common/service"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// Test for the endpoint GET /restaurantGroups/{id}
func TestGetRestaurantGroupById(t *testing.T) {

	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)
	mockRestaurantGroup := model.RestaurantGroup{ID: "1", Name: "Test RestaurantGroup", Description: "This is a test restaurantGroup"}

	singleResultMock := mongo.NewSingleResultFromDocument(mockRestaurantGroup, nil, nil)
	collMock.On("FindOne", mock.Anything).Return(singleResultMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.RestaurantGroup](collMock)
	repo := baserepo.NewLoggingRepository[model.RestaurantGroup](repoWithoutLog)

	service := baseservice.NewService[model.RestaurantGroup](repo)
	handler := basehandler.NewCommonHandler[model.RestaurantGroup](service)

	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.RestaurantGroup](r, &handler, "restaurantGroups")

	// 3. Create a GET request
	req := httptest.NewRequest(http.MethodGet, "/restaurantGroups/1", nil)
	rec := httptest.NewRecorder()

	c := r.NewContext(req, rec)
	c.SetPath("/restaurantGroups/:id")
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

	// TODO: Additional checks for the returned restaurantGroup

	collMock.AssertNumberOfCalls(t, "FindOne", 1)
}
