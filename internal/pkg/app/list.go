package app

import (
	"encoding/json"
	"net/http"
	"recipes/internal/pkg/common"
	"recipes/internal/pkg/service"
	"strconv"
	"strings"
)

// ResponseObject is the json response
type ResponseObject struct {
	Recipes []common.Recipe        `json:"recipes"`
	List    map[string]*Ingredient `json:"list"`
}

// Ingredient is a subset of shopping List
type Ingredient struct {
	Unit     string  `json:"unit"`
	Quantity float64 `json:"quantity"`
}

// CombineIngredients creates combined values/units
func CombineIngredients(r []common.Recipe) map[string]*Ingredient {
	parentMeasurement := map[string]string{
		"gram":       "kilogram",
		"millilitre": "litre",
	}

	childMeasurement := map[string]string{
		"kilogram": "gram",
		"litre":    "millilitre",
	}

	ingredientList := make(map[string]*Ingredient)

	for _, recipe := range r {
		for _, ingredient := range recipe.Ingredients {
			existingIngredient, exists := ingredientList[ingredient.Name]
			if q, err := strconv.ParseFloat(ingredient.Quantity, 64); err == nil {
				// convert to childMeasurement if needed
				if newUnit, isParentMeasurement := childMeasurement[ingredient.Unit]; isParentMeasurement {
					q = q * 1000
					ingredient.Unit = newUnit
				}
				if exists {
					existingIngredient.Quantity = existingIngredient.Quantity + q
				} else {
					newIngredient := Ingredient{
						Unit:     ingredient.Unit,
						Quantity: q,
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
		if newUnit, exists := parentMeasurement[value.Unit]; exists {
			ingredientList[key].Unit = newUnit
			ingredientList[key].Quantity = ingredientList[key].Quantity / 1000
		}
	}

	return ingredientList
}

func (a *App) getListHandler(w http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	recipeIDs := strings.Split(qs.Get("recipes"), ",")

	response := &ResponseObject{}

	for i := 0; i < len(recipeIDs); i++ {
		id, err := strconv.Atoi(recipeIDs[i])
		if err == nil {
			// do something
		}
		recipe, err := service.GetRecipeByID(id, a.db)
		if err != nil {
			// do something
		}
		response.Recipes = append(response.Recipes, *recipe)
	}

	response.List = CombineIngredients(response.Recipes)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}

}
