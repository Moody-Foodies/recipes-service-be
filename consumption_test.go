package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"encoding/json"
	"example.com/recipes-service-be/recipes"
	"github.com/stretchr/testify/require"
)

func TestMakeRequest(t *testing.T) {
	mockReturn, err := os.ReadFile("recipes/testdata/magnesium_fixture_data.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	mockResponse, err := os.ReadFile("recipes/testdata/edamam_response_data.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	testCase := struct {
		name string
		retCode int
		retBody []byte
		wantBody string
	}{
		name: "Response Test",
		retCode: 200,
		retBody: mockResponse,
		wantBody: string(mockReturn),
	}

	t.Run(testCase.name, func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			w.WriteHeader(testCase.retCode)
			_, err := w.Write(testCase.retBody)
			require.NoError(t, err)
		}))
		defer testServer.Close()
		
		gotBody, _ := json.Marshal(recipes.MakeRequest("Magnesium", "600", testServer.URL))
		require.EqualValues(t, testCase.wantBody, string(gotBody))
	})

}