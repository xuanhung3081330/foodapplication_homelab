package implementations

import (
	"foodapplication/models"
	"foodapplication/repositories/interfaces"
	"foodapplication/config"
	"time"
)

type FoodRepository struct {}

func NewFoodRepository() interfaces.FoodRepository{
	return &FoodRepository{}
}


func (foodRepo *FoodRepository) CreateFood(food *models.Food) (int, error){
	var id int
	err := config.DB.QueryRow(
		"INSERT INTO foods (description, category, price, created_at, created_by) VALUES ($1, $2, $3, $4, $5) RETURNING id", food.Description, food.Category, food.Price, time.Now(), "admin-hungnx").Scan(&id)
	if err != nil{
		return 0, err
	}

	return id, nil
}
