package service

import (
	"database/sql"
	"log"
)

// Recipe is a lightweight recipe type w/o ingredients
type Recipe struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// GetAllRecipes returns all recipes in the recipe table
func GetAllRecipes(db *sql.DB, userID string) ([]Recipe, error) {
	accountID, err := GetAccountID(db, userID)
	recipesQuery := `
		SELECT id, name FROM recipe
			WHERE account_id = ?
			ORDER BY lower(recipe.name);
	`
	results, err := db.Query(recipesQuery, accountID)

	if err != nil {
		log.Println("Error querying recipes")
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
