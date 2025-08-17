package implementations

import (
	"foodapplication/models"
	repoInterfaces "foodapplication/repositories/interfaces"
	svcInterfaces "foodapplication/services/interfaces"
)

type FoodService struct{
	repo repoInterfaces.FoodRepository
}

func NewFoodService(repo repoInterfaces.FoodRepository) svcInterfaces.FoodService{
	return &FoodService{repo: repo}
}

func (foodSvc *FoodService) CreateFood(food *models.Food) (int, error){
	return foodSvc.repo.CreateFood(food)
}
