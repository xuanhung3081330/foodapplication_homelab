package interfaces

import (
	"foodapplication/models"
)

type FoodRepository interface {
	CreateFood(food *models.Food) (int, error)
}
