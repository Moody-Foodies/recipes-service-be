package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
	router := gin.Default()
	return router
}

func TestMakeRequest(t *testing.T) {
	mockResponse, err := os.ReadFile("testdata/magnesium_fixture.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	r := SetUpRouter()
	r.GET("/recipes", func(c *gin.Context) {
		nutrient := c.Query("nutrient")
		cookTime := c.Query("cook_time")
		MakeRequest(c, nutrient, cookTime)
	})

	req, _ := http.NewRequest("GET", "/recipes?nutrient=Magnesium&cook_time=600", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, string(mockResponse), string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}