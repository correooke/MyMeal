package handler

import (
	"bytes"
	"resto-service/api/model"
	"resto-service/internal/app/resto/handler/mocks"
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

// Test for the endpoint PUT /restos/{id}
func TestUpdateResto(t *testing.T) {
	// 1. Setup in-memory database and mocks
	collMock := new(mocks.CollectionMock)

	updatedResto := model.Resto{ID: "1", Name: "Updated Resto", Description: "This is an updated resto"}

	ur := &mongo.UpdateResult{
		MatchedCount:  1,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    "1",
	}

	collMock.On("UpdateOne", mock.Anything).Return(ur, nil)

	repoWithoutLog := baserepo.NewMongoRepository[model.Resto](collMock)
	repo := baserepo.NewLoggingRepository[model.Resto](repoWithoutLog)

	service := baseservice.NewService[model.Resto](repo)
	handler := basehandler.NewCommonHandler[model.Resto](service)

	r := baserouter.NewRouter()
	baserouter.AddCRUDRoutes[model.Resto](r, &handler, "restos")

	// Prepare data for PUT request
	jsonData, _ := json.Marshal(updatedResto)

	// 3. Create a PUT request
	req := httptest.NewRequest(http.MethodPut, "/restos/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	c := r.NewContext(req, rr)
	c.SetPath("/restos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// 4. Call the corresponding handler
	handler.Update(c)

	// 5. Check the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Additional checks for the updated resto

	collMock.AssertNumberOfCalls(t, "UpdateOne", 1)

}
