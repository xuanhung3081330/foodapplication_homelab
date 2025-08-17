package interfaces

import "foodapplication/models"

type FoodService interface {
	CreateFood(food *models.Food) (int, error)
}
