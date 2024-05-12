package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func MakeRequest(c *gin.Context, nutrient string, cookTime string) {
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?instructionsRequired=true&addRecipeInstructions=true&min%s=10&apiKey=2a5896d0a98840049a49c52f14121a06&addRecipeInformation=true&maxReadyTime=%s&fillIngredients=true", nutrient, cookTime)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	var data []map[string]any

	for _, recipe := range responseObject.Recipes {
		data = append(data, map[string]any{
			"id":           recipe.ID,
			"title":        recipe.Title,
			"cook_time":    recipe.CookTime,
			"image":        recipe.Image,
			"ingredients":  recipe.Ingredients,
			"instructions": recipe.Instructions,
		})
	}
	var payload = Payload{
		Data: data,
	}

	c.JSON(http.StatusOK, payload)
}

func main() {
	router := gin.Default()
	router.GET("/recipes", func(c *gin.Context) {
		nutrient := c.Query("nutrient")
		cookTime := c.Query("cook_time")
		MakeRequest(c, nutrient, cookTime)
	})
	router.Run("localhost:8080")
}

type Payload struct {
	Data any `json:"data"`
}

type Response struct {
	Recipes []Recipe `json:"results"`
}

type Recipe struct {
	ID           int            `json:"id"`
	Title        string         `json:"title"`
	CookTime     int            `json:"readyInMinutes"`
	Image        string         `json:"image"`
	Ingredients  []Ingredient   `json:"extendedIngredients"`
	Instructions []Instructions `json:"analyzedInstructions"`
}

type Ingredient struct {
	Name string `json:"original"`
}

type Instructions struct {
	Steps []Step `json:"steps"`
}

type Step struct {
	Step string `json:"step"`
}
