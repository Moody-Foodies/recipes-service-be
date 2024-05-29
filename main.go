package main

import (
	"example.com/recipes-service-be/recipes"
	"github.com/gin-gonic/gin"
	"fmt"
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
		recipes.MakeRequest(c, nutrient, cookTime)
	})

	router.Run()
}
