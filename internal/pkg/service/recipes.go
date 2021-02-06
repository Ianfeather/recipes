package service

import (
	"database/sql"
	"fmt"
)

// Recipe is a lightweight recipe type w/o ingredients
type Recipe struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// GetAllRecipes returns all recipes in the recipe table
func GetAllRecipes(db *sql.DB, userID int) ([]Recipe, error) {
	recipesQuery := `
		SELECT recipe.id as id, recipe.name as name FROM recipe_user
			INNER JOIN recipe on recipe_user.recipe_id = recipe.id
			WHERE recipe_user.user_id = ?
			ORDER BY lower(recipe.name);
	`
	results, err := db.Query(recipesQuery, userID)

	if err != nil {
		fmt.Println("Error querying recipes")
		return nil, err
	}

	recipes := make([]Recipe, 0)

	for results.Next() {
		r := Recipe{}
		err = results.Scan(&r.ID, &r.Name)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, r)
	}
	return recipes, nil
}
