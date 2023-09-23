package handler

import (
	"bytes"
	"dish-service/api/model"
	"dish-service/internal/app/dish/handler/mocks"
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

// Test for the endpoint PUT /dishes/{id}
func TestUpdateDish(t *testing.T) {
	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)

	updatedDish := model.Dish{ID: "1", Name: "Updated Dish", Description: "This is an updated dish"}

	ur := &mongo.UpdateResult{
		MatchedCount:  1,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    "1",
	}

	collMock.On("UpdateOne", mock.Anything).Return(ur, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Dish](collMock)
	repo := baserepo.NewLoggingRepository[model.Dish](repoWithoutLog)

	service := baseservice.NewService[model.Dish](repo)
	handler := basehandler.NewCommonHandler[model.Dish](service)

	r := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Dish](r, &handler, "dishes")

	// Prepare data for PUT request
	jsonData, _ := json.Marshal(updatedDish)

	// 3. Create a PUT request
	req := httptest.NewRequest(http.MethodPut, "/dishes/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	c := r.NewContext(req, rr)
	c.SetPath("/dishes/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// 4. Call the corresponding handler
	handler.Update(c)

	// 5. Check the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Additional checks for the updated dish

	collMock.AssertNumberOfCalls(t, "UpdateOne", 1)

}
