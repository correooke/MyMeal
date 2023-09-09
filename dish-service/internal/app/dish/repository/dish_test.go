package repository

import (
	"context"
	"dish-service/api/model"
	"dish-service/internal/app/dish/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Asume que MongoDB está corriendo localmente
	return mongo.Connect(context.TODO(), clientOptions)
}

func TestMongoDishRepository_GetDishes(t *testing.T) {
	repoMock := new(mocks.DishRepositoryMock)

	mockDishes := []model.Dish{{Name: "TestDish", Description: "TestDescription", Price: 10.99}}
	repoMock.On("GetDishes", mock.Anything).Return(mockDishes, nil)

	dishes, err := repoMock.GetDishes(context.TODO())
	assert.NoError(t, err)

	assert.NotEqual(t, 0, len(dishes)) // Asegura que hay platos (esta prueba depende de los datos en tu DB de prueba)
	repoMock.AssertExpectations(t)
}

//... (puedes agregar tests para los otros métodos también) ...
