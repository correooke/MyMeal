package handler

import (
	"restaurantGroup-service/api/model"
	"restaurantGroup-service/internal/app/restaurantGroup/handler/mocks"
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

// Test for the endpoint GET /restaurantGroups
func TestGetAllCategories(t *testing.T) {

	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)
	mockCategories := []model.RestaurantGroup{
		{ID: "1", Name: "Test RestaurantGroup 1", Description: "This is a test restaurantGroup 1"},
		{ID: "2", Name: "Test RestaurantGroup 2", Description: "This is a test restaurantGroup 2"},
	}
	var docs []interface{}
	for _, restaurantGroup := range mockCategories {
		docs = append(docs, restaurantGroup)
	}
	cursorMock, _ := mongo.NewCursorFromDocuments(docs, nil, nil)
	collMock.On("Find", mock.Anything).Return(cursorMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.RestaurantGroup](collMock)
	repo := baserepo.NewLoggingRepository[model.RestaurantGroup](repoWithoutLog)

	service := baseservice.NewService[model.RestaurantGroup](repo)
	handler := basehandler.NewCommonHandler[model.RestaurantGroup](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.RestaurantGroup](r2, &handler, "restaurantGroups")

	// 3. Create a GET request
	req := httptest.NewRequest(http.MethodGet, "/restaurantGroups", nil)
	rr := httptest.NewRecorder()

	c := r2.NewContext(req, rr)
	handler.GetAll(c)

	// 5. Check the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Additional checks for the returned restaurantGroups

	collMock.AssertNumberOfCalls(t, "Find", 1)

}
