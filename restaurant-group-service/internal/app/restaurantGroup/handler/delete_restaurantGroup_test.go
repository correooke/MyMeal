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

// Test for the endpoint DELETE /restaurantGroups/{id}
func TestDeleteRestaurantGroup(t *testing.T) {
	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)

	deleteResultMock := &mongo.DeleteResult{
		DeletedCount: 1,
	}

	collMock.On("DeleteOne", mock.Anything).Return(deleteResultMock, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.RestaurantGroup](collMock)
	repo := baserepo.NewLoggingRepository[model.RestaurantGroup](repoWithoutLog)

	service := baseservice.NewService[model.RestaurantGroup](repo)
	handler := basehandler.NewCommonHandler[model.RestaurantGroup](service)

	r2 := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.RestaurantGroup](r2, &handler, "restaurantGroups")

	// 3. Create a DELETE request
	req := httptest.NewRequest(http.MethodDelete, "/restaurantGroups/1", nil)
	rr := httptest.NewRecorder()

	c := r2.NewContext(req, rr)
	c.SetPath("/restaurantGroups/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler.Delete(c)

	// 5. Check the response
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// TODO: Additional checks for the deletion result

	collMock.AssertNumberOfCalls(t, "DeleteOne", 1)
}
