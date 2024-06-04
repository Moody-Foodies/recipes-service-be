# BrainFood Recipe Micro Service
### Summary:
This application acts as a micro service for the BrainFood BE Gateway application. It's responsibility is to make HTTP Requests to the Edamam Recipe Search API endpoint with specific query parameters of nutrients required to be present in recipes, based on information provided in the requests from the internal, backend gateway. It then takes the response that it receives from the external api, restructures the data, and returns it to the BrainFood gateway application to be served to our frontend application. It utilizes a web scraper built with the go-colly package in order to scrape recipe instructions from the webpages which recipes are lifted from and restructure them into its response.

### Versioning
- go v1.22.3 darwin/arm64

### Installation
- Clone this repository to your local machine `git clone <repository url>`
- CD into the repository `cd recipes-service-be`
- Install Golang on your computer by running the following command in your terminal if you have brew installed `brew install go`
    - If you do not have brew installed, you can visit this link to install Go directly from your browser `https://go.dev/doc/install`
- Create a `.env` file in the root directory and add in your free API Key and App ID from the Edamam Recipe Search API as per their instructions
    - The file should look like the below, with your API Key and App ID replacing the placeholders
```
API_KEY=<your-api-key-here>
APP_ID=<your-app-id-here>
```
- In your terminal, run the following command to install all dependencies and start your local host server `go run main.go`
- Begin making calls to the endpoint listed below with the base URL of `http://localhost:8080`

### Endpoint
```
GET "/recipes?nutrient={nutrient_name}&cook_time={cook_time_in_minutes}"
```

Response:
```
{
    "data": [
        {
            "cook_time": 55,
            "id": 715415,
            "image": "https://fake-image-url.jpg",
            "ingredients": [ "additional toppings: diced avocado, micro greens, chopped basil", "3 medium carrots, peeled and diced", "3 celery stalks, diced", "2 cups fully-cooked chicken breast, shredded (may be omitted for a vegetarian version)", "½ cup flat leaf Italian parsley, chopped (plus extra for garnish)", "6 cloves of garlic, finely minced", "2 tablespoons olive oil", "28 ounce-can plum tomatoes, drained and rinsed, chopped", "2 cups dried red lentils, rinsed", "salt and black pepper, to taste", "1 large turnip, peeled and diced", "8 cups vegetable stock", "1 medium yellow onion, diced"],
            "instructions": ["To a large dutch oven or soup pot, heat the olive oil over medium heat.", "Add the onion, carrots and celery and cook for 8-10 minutes or until tender, stirring occasionally.", "Add the garlic and cook for an additional 2 minutes, or until fragrant. Season conservatively with a pinch of salt and black pepper.To the pot, add the tomatoes, turnip and red lentils. Stir to combine. Stir in the vegetable stock and increase the heat on the stove to high. Bring the soup to a boil and then reduce to a simmer. Simmer for 20 minutes or until the turnips are tender and the lentils are cooked through.", "Add the chicken breast and parsley. Cook for an additional 5 minutes. Adjust seasoning to taste.", "Serve the soup immediately garnished with fresh parsley and any additional toppings. Enjoy!"]
            "title": "Red Lentil Soup with Chicken and Turnips",
            "description": "A magnesium rich recipe to help boost your mood"
        },
        ....
    ]
}        
```