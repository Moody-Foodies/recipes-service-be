package recipes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"example.com/recipes-service-be/scraper"
)

func MakeRequest(nutrient string, cookTime string, apiURL string) (Payload) {
	var nutrientNew string

	if nutrient == "Folic Acid" {
		nutrientNew = "%5BFOLAC%5D"
	}
	if nutrient == "Magnesium" {
		nutrientNew = "%5BMG%5D"
	}
	if nutrient == "Fiber" {
		nutrientNew = "%5BFIBTG%5D"
	}
	if nutrient == "Vitamin B12" {
		nutrientNew = "%5BVITB12%5D"
	}
	if nutrient == "Vitamin D" {
		nutrientNew = "%5BVITD%5D"
	}

	apiKey := os.Getenv("API_KEY")
	appID := os.Getenv("APP_ID")

	url := fmt.Sprintf("%s?type=public&app_id=%s&app_key=%s&mealType=Lunch&time=%s&imageSize=REGULAR&nutrients%s=10&dishType=main_course", apiURL, appID, apiKey, cookTime, nutrientNew)

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
	var instructions []string

	for _, recipe := range responseObject.Recipes {
		instructions = scraper.ScrapeInstructions(recipe.RecipeInfo.Url)

		if instructions == nil {
			instructions = append(instructions, recipe.RecipeInfo.Url)
		}

		data = append(data, map[string]any{
			"id":           1,
			"title":        recipe.RecipeInfo.Title,
			"cook_time":    recipe.RecipeInfo.CookTime,
			"image":        recipe.RecipeInfo.Images.Large.Url,
			"description":  fmt.Sprintf("%s rich recipe to help your current mood", nutrient),
			"ingredients":  recipe.RecipeInfo.Ingredients,
			"instructions": instructions,
		})
	}

	var payload = Payload{
		Data: data,
	}

	return payload
}

type Payload struct {
	Data any `json:"data"`
}

type Response struct {
	Recipes []Recipe `json:"hits"`
}

type Recipe struct {
	RecipeInfo RecipeInfo `json:"recipe"`
}
type RecipeInfo struct {
	ID          int      `json:"id"`
	Title       string   `json:"label"`
	Images      Images   `json:"images"`
	CookTime    any      `json:"totalTime"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Ingredients []string `json:"ingredientLines"`
}

type Images struct {
	Large Url `json:"LARGE"`
}

type Url struct {
	Url string `json:"url"`
}
