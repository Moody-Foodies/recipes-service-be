package main

import (
	"fmt"
	"net/http"
	"example.com/recipes-service-be/recipes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading.env file")
	}

	router := gin.Default()
	router.GET("/recipes", func(c *gin.Context) {
		nutrient := c.Query("nutrient")
		cookTime := c.Query("cook_time")
		c.JSON(http.StatusOK, recipes.MakeRequest(nutrient, cookTime, "https://api.edamam.com/api/recipes/v2"))
	})

	router.Run()
}
