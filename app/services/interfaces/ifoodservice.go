package interfaces

import "foodapplication/models"

type IFoodService interface {
	CreateFood(food *models.Food) (int, error)
}
