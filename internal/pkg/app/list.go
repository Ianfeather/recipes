package app

import (
	"net/http"
	"recipes/internal/pkg/common"
	"recipes/internal/pkg/service"
	"strings"

	"github.com/gorilla/mux"
)

// RecipeList is a list of recipes
type RecipeList struct {
	recipes []common.Recipe
}

// GetListHandler returns a shopping list
func (a *App) GetListHandler(w http.ResponseWriter, req *http.Request) {
	recipes := mux.Vars(req)["recipes"]
	recipeIDs := strings.Split(recipes, ",")

	recipeList := &RecipeList{}

	for i := 0; i < len(recipeIDs); i++ {
		recipe, err := service.GetRecipeByID(recipeIDs[i], a.db)
		if err != nil {
			// do something
		}
		recipeList.recipes = append(recipeList.recipes, *recipe)
	}

}
