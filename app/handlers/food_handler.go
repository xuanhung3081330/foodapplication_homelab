package handlers

import (
	"foodapplication/models"
	svcInterfaces "foodapplication/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	Service svcInterfaces.IFoodService
}

func NewFoodHandler(svc svcInterfaces.IFoodService) *FoodHandler{
	return &FoodHandler{Service: svc}
}

func (handler *FoodHandler) CreateFood(c *gin.Context){
	var food models.Food
	if err := c.BindJSON(&food); err != nil {
		// Binding failed (bad request, invalid JSON, etc.)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := handler.Service.CreateFood(&food)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
