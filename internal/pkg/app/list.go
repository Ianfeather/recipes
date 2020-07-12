package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recipes/internal/pkg/common"
	"recipes/internal/pkg/service"
	"strconv"
)

// ResponseObject is the json response
type ResponseObject struct {
	Recipes 		[]common.Recipe        `json:"recipes"`
	Ingredients map[string]*Ingredient `json:"ingredients"`
	Extras    	map[string]*Ingredient `json:"extras"`
}

// Ingredient is a subset of shopping List
type Ingredient struct {
	Unit      string  `json:"unit"`
	Quantity  float64 `json:"quantity"`
	IsChecked bool    `json:"isChecked"`
}

// CombineIngredients creates combined values/units
func CombineIngredients(r []common.Recipe) map[string]*Ingredient {
	parentUnit := map[string]string{
		"gram":       "kilogram",
		"millilitre": "litre",
	}
	childUnit := map[string]string{
		"kilogram": "gram",
		"litre":    "millilitre",
	}

	ingredientList := make(map[string]*Ingredient)
	for _, recipe := range r {
		for _, ingredient := range recipe.Ingredients {
			if q, err := strconv.ParseFloat(ingredient.Quantity, 64); err == nil {
				if childUnit, isParentUnit := childUnit[ingredient.Unit]; isParentUnit {
					q = q * 1000
					ingredient.Unit = childUnit
				}
				if existingIngredient, exists := ingredientList[ingredient.Name]; exists {
					existingIngredient.Quantity = existingIngredient.Quantity + q
				} else {
					newIngredient := Ingredient{
						Unit:     ingredient.Unit,
						Quantity: q,
						IsChecked: false
					}
					ingredientList[ingredient.Name] = &newIngredient
				}
			}
		}
	}

	for key, value := range ingredientList {
		if value.Quantity < 1000 {
			continue
		}
		if parentUnit, exists := parentUnit[value.Unit]; exists {
			ingredientList[key].Unit = parentUnit
			ingredientList[key].Quantity = ingredientList[key].Quantity / 1000
		}
	}

	return ingredientList
}

func (a *App) postListHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: Identity
	userID := 1

	recipeIDs := make([]string, 0)
	err := json.NewDecoder(req.Body).Decode(&recipeIDs)

	if err != nil {
		w.Write([]byte("Error decoding json body"))
	}

	response := &ResponseObject{}

	for i := 0; i < len(recipeIDs); i++ {
		id, err := strconv.Atoi(recipeIDs[i])
		if err != nil {
			fmt.Println(err)
		}
		recipe, err := service.GetRecipeByID(id, a.db)
		if err != nil {
			fmt.Println(err)
		}
		response.Recipes = append(response.Recipes, *recipe)
	}

	combinedIngredients := CombineIngredients(response.Recipes)

	service.removeIngredientListItems(userID, a.db)

	// Replace all the items
	service.addIngredientListItems(userID, combinedIngredients, a.db)

	// Return the new items + extras
	response.Ingredients = combinedIngredients
	response.Extras = service.getExtraListItems(userID, a.db)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}

}
