package recipes

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

	apiKey := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?instructionsRequired=true&addRecipeInstructions=true&min%s=10&apiKey=%s&addRecipeInformation=true&maxReadyTime=%s&fillIngredients=true", nutrient, apiKey, cookTime)

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
			"description":  "Placeholder Description",
			"ingredients":  recipe.Ingredients,
			"instructions": recipe.Instructions,
		})
	}

	var payload = Payload{
		Data: data,
	}

	c.JSON(http.StatusOK, payload)
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
	Description  string         `json:"description"`
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
