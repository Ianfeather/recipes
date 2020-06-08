package service

import (
	"database/sql"
	"fmt"
	"recipes/internal/pkg/common"
)

// GetRecipe fetches a recipe from the database
func GetRecipe(slug string, db *sql.DB) (r *common.Recipe, e error) {
	recipe := &common.Recipe{Ingredients: []common.Ingredient{}}
	recipeQuery := "SELECT id, name FROM recipe where slug='%s'"

	err := db.QueryRow(fmt.Sprintf(recipeQuery, slug)).Scan(&recipe.ID, &recipe.Name)

	if err != nil {
		return nil, err
	}

	ingredientQuery := "SELECT ingredient.name as name, unit.name as unit, quantity FROM part INNER JOIN ingredient on ingredient_id = ingredient.id INNER JOIN unit on unit_id = unit.id WHERE recipe_id = %d;"
	results, err := db.Query(fmt.Sprintf(ingredientQuery, recipe.ID))

	for results.Next() {
		ingredient := common.Ingredient{}
		err = results.Scan(&ingredient.Name, &ingredient.Unit, &ingredient.Quantity)
		if err != nil {
			return nil, err
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}
	return recipe, nil
}
