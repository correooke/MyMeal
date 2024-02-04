
package main

import (
	"context"
	"encoding/json"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/correooke/MyMeal/common/model"
	"github.com/correooke/MyMeal/common/service"
	"dto"
)

func MapDishToDTO(dish model.Dish) dto.DishDTO {
	return dto.DishDTO{
		ID:          dish.ID,
		Name:        dish.Name,
		Description: dish.Description,
		Image:       dish.Image,
		Price:       dish.Price,
	}
}

func MapCategoryToDTO(cat model.Category, dishService service.CommonService[model.Dish]) dto.CategoryDTO {
	dishes := []dto.DishDTO{}
	for _, dishID := range cat.Dishes {
		dish, err := dishService.GetByID(context.TODO(), dishID)
		if err != nil {
			log.Printf("Error retrieving dish with ID %s: %v", dishID, err)
			continue
		}
		dishes = append(dishes, MapDishToDTO(dish.(model.Dish)))
	}
	return dto.CategoryDTO{
		ID:          cat.ID,
		Name:        cat.Name,
		Description: cat.Description,
		Image:       cat.Image,
		Dishes:      dishes,
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	restoID := request.PathParameters["id"]
	// Assuming you have the repo implementations, you can initialize the services like this:
	// restoService := service.NewService[model.Resto](restoRepo)
	// categoryService := service.NewService[model.Category](categoryRepo)
	// dishService := service.NewService[model.Dish](dishRepo)
	resto, err := restoService.GetByID(context.TODO(), restoID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Resto not found",
		}, nil
	}
	categories := []dto.CategoryDTO{}
	for _, menu := range resto.(model.Resto).Menues {
		for _, categoryID := range menu {
			cat, err := categoryService.GetByID(context.TODO(), categoryID)
			if err != nil {
				log.Printf("Error retrieving category with ID %s: %v", categoryID, err)
				continue
			}
			categories = append(categories, MapCategoryToDTO(cat.(model.Category), dishService))
		}
	}
	restoDTO := dto.RestoDTO{
		ID:            resto.(model.Resto).ID,
		Name:          resto.(model.Resto).Name,
		Description:   resto.(model.Resto).Description,
		Address:       resto.(model.Resto).Address,
		Configuration: resto.(model.Resto).Configuration,
		MainImage:     resto.(model.Resto).MainImage,
		DesignType:    resto.(model.Resto).DesignType,
		Currency:      resto.(model.Resto).Currency,
		Owner:         resto.(model.Resto).Owner,
		Menues:        resto.(model.Resto).Menues,
		Categories:    categories,
	}
	// Return the mapped DTO
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       json.Marshal(restoDTO),
	}, nil
}

func main() {
	lambda.Start(handler)
}
