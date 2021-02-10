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
func GetAllRecipes(db *sql.DB, userID string) ([]Recipe, error) {
	recipesQuery := `
		SELECT id, name FROM recipe
			WHERE user_id = ?
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
