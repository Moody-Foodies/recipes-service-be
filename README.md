# BrainFood Recipe Micro Service
### Summary:
This application acts as a micro service for the BrainFood BE Gateway application. It's responsibility is to make HTTP requests to the external Spoonacular API to retrieve full recipes based on the criteria of nutrients that should be included in the recipe and the maximum amount of time that the recipe should take to prepare and cook. It then restructures the needed data to respond to the BE gateway application to utilize for the FE of the app.

### Versioning
- go v1.22.3 darwin/arm64

### Endpoints
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
            "image": "https://img.spoonacular.com/recipes/715415-312x231.jpg",
            "ingredients": [
                {
                    "original": "additional toppings: diced avocado, micro greens, chopped basil)"
                },
                {
                    "original": "3 medium carrots, peeled and diced"
                },
                {
                    "original": "3 celery stalks, diced"
                },
                {
                    "original": "2 cups fully-cooked chicken breast, shredded (may be omitted for a vegetarian version)"
                },
                {
                    "original": "½ cup flat leaf Italian parsley, chopped (plus extra for garnish)"
                },
                {
                    "original": "6 cloves of garlic, finely minced"
                },
                {
                    "original": "2 tablespoons olive oil"
                },
                {
                    "original": "28 ounce-can plum tomatoes, drained and rinsed, chopped"
                },
                {
                    "original": "2 cups dried red lentils, rinsed"
                },
                {
                    "original": "salt and black pepper, to taste"
                },
                {
                    "original": "1 large turnip, peeled and diced"
                },
                {
                    "original": "8 cups vegetable stock"
                },
                {
                    "original": "1 medium yellow onion, diced"
                }
            ],
            "instructions": [
                {
                    "steps": [
                        {
                            "step": "To a large dutch oven or soup pot, heat the olive oil over medium heat."
                        },
                        {
                            "step": "Add the onion, carrots and celery and cook for 8-10 minutes or until tender, stirring occasionally."
                        },
                        {
                            "step": "Add the garlic and cook for an additional 2 minutes, or until fragrant. Season conservatively with a pinch of salt and black pepper.To the pot, add the tomatoes, turnip and red lentils. Stir to combine. Stir in the vegetable stock and increase the heat on the stove to high. Bring the soup to a boil and then reduce to a simmer. Simmer for 20 minutes or until the turnips are tender and the lentils are cooked through."
                        },
                        {
                            "step": "Add the chicken breast and parsley. Cook for an additional 5 minutes. Adjust seasoning to taste."
                        },
                        {
                            "step": "Serve the soup immediately garnished with fresh parsley and any additional toppings. Enjoy!"
                        }
                    ]
                }
            ],
            "title": "Red Lentil Soup with Chicken and Turnips"
        },
        {
            "cook_time": 20,
            "id": 716406,
            "image": "https://img.spoonacular.com/recipes/716406-312x231.jpg",
            "ingredients": [
                {
                    "original": "1 bag of frozen organic asparagus (preferably thawed)"
                },
                {
                    "original": "1T EVOO (extra virgin olive oil)"
                },
                {
                    "original": "a couple of garlic cloves"
                },
                {
                    "original": "1/2 onion"
                },
                {
                    "original": "2-3c of frozen organic peas"
                },
                {
                    "original": "1 box low-sodium vegetable broth"
                }
            ],
            "instructions": [
                {
                    "steps": [
                        {
                            "step": "Chop the garlic and onions."
                        },
                        {
                            "step": "Saute the onions in the EVOO, adding the garlic after a couple of minutes; cook until the onions are translucent."
                        },
                        {
                            "step": "Add the whole bag of asparagus and cover everything with the broth. Season with salt and pepper and a pinch of red pepper flakes, if using.Simmer until the asparagus is bright green and tender (if you've thawed the asparagus it will only take a couple of minutes). Turn off the heat and puree using an immersion blender."
                        },
                        {
                            "step": "Add peas (the heat of the soup will quickly thaw them) and puree until smooth; add more until it reaches the thickness you like.Top with chives and a small dollop of creme fraiche or sour cream or greek yogurt."
                        }
                    ]
                }
            ],
            "title": "Asparagus and Pea Soup: Real Convenience Food"
        },
        ....
```