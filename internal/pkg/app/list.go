package app

import (
	"encoding/json"
	"net/http"
	"recipes/internal/pkg/common"
	"recipes/internal/pkg/service"
	"strconv"
	"strings"
)

// RecipeList is a list of recipes
type RecipeList struct {
	Recipes []common.Recipe `json:"recipes"`
}

func (a *App) getListHandler(w http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	recipeIDs := strings.Split(qs.Get("recipes"), ",")

	recipeList := &RecipeList{}

	for i := 0; i < len(recipeIDs); i++ {
		id, err := strconv.Atoi(recipeIDs[i])
		if err == nil {
			// do something
		}
		recipe, err := service.GetRecipeByID(id, a.db)
		if err != nil {
			// do something
		}
		recipeList.Recipes = append(recipeList.Recipes, *recipe)
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(recipeList)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}

}
