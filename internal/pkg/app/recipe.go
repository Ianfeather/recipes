package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Ingredient contains ingredient fields
type Ingredient struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Recipe contains recipe fields
type Recipe struct {
	Name        string       `json:"name"`
	ID          int          `json:"id"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func slugify(s string) string {
	slug := strings.ReplaceAll(s, " ", "-")
	slugLength := Min(60, len(slug))
	return strings.ToLower(slug[0:slugLength])
}

func (a *App) recipeHandler(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)["slug"]

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

	ingredientQuery := "SELECT ingredient_id, name FROM component INNER JOIN ingredient ON ingredient_id = ingredient.id where recipe_id = %d;"
	results, err := a.db.Query(fmt.Sprintf(ingredientQuery, recipe.ID))

	for results.Next() {
		ingredient := Ingredient{}
		err = results.Scan(&ingredient.ID, &ingredient.Name)
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

	_, err = a.db.Query(fmt.Sprintf("INSERT INTO recipe (name, slug) VALUES ('%s', '%s')", recipe.Name, slugify(recipe.Name)))

	if err != nil {
		fmt.Println("could not insert recipe")
		fmt.Println(err.Error())
	}

	fmt.Printf("Stored %s", recipe.Name)
}
