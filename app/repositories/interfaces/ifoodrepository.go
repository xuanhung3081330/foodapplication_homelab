package interfaces

import (
	"foodapplication/models"
)

type IFoodRepository interface {
	CreateFood(food *models.Food) (int, error)
}
