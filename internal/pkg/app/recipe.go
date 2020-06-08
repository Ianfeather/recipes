package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"recipes/internal/pkg/common"

	"github.com/gorilla/mux"
)

// Ingredient contains ingredient fields
type Ingredient struct {
	Name     string `json:"name"`
	Unit     string `json:"unit"`
	Quantity string `json:"quantity"`
}

// Recipe contains recipe fields
type Recipe struct {
	Name        string       `json:"name"`
	ID          int          `json:"id"`
	Ingredients []Ingredient `json:"ingredients"`
}

func (a *App) recipeHandler(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)["slug"]

	// TODO: turn it into a single query
	recipe := Recipe{Ingredients: []Ingredient{}}
	recipeQuery := "SELECT id, name FROM recipe where slug='%s'"

	err := a.db.QueryRow(fmt.Sprintf(recipeQuery, slug)).Scan(&recipe.ID, &recipe.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Recipe not found", 404)
			return
		}
		fmt.Sprintln("Failed to parse recipe from db")
		http.Error(w, err.Error(), 500)
		return
	}

	ingredientQuery := "SELECT ingredient.name as name, unit.name as unit, quantity FROM part INNER JOIN ingredient on ingredient_id = ingredient.id INNER JOIN unit on unit_id = unit.id WHERE recipe_id = %d;"
	results, err := a.db.Query(fmt.Sprintf(ingredientQuery, recipe.ID))

	for results.Next() {
		ingredient := Ingredient{}
		err = results.Scan(&ingredient.Name, &ingredient.Unit, &ingredient.Quantity)
		if err != nil {
			fmt.Sprintln("Failed to parse ingredient from db")
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(recipe)

	if err != nil {
		w.Write([]byte("Error encoding json"))
	}
}

func (a *App) addRecipeHandler(w http.ResponseWriter, req *http.Request) {
	recipe := Recipe{}
	err := json.NewDecoder(req.Body).Decode(&recipe)

	if err != nil {
		w.Write([]byte("Error decoding json body"))
	}

	_, err = a.db.Query(fmt.Sprintf("INSERT INTO recipe (name, slug) VALUES ('%s', '%s')", recipe.Name, common.Slugify(recipe.Name)))

	if err != nil {
		fmt.Println("could not insert recipe")
		fmt.Println(err.Error())
	}

	fmt.Printf("Stored %s", recipe.Name)
}
