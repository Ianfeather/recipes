package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Recipe contains recipe fields
type Recipe struct {
	Name        string
	Ingredients int
}

var myRecipe = Recipe{"Braised Cabbage", 6}

func (a *App) recipeHandler(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(myRecipe)
	if err != nil {
		w.Write([]byte("Error encoding json"))
	}
}

func (a *App) addRecipeHandler(w http.ResponseWriter, req *http.Request) {
	err := json.NewDecoder(req.Body).Decode(&myRecipe)

	if err != nil {
		w.Write([]byte("Error decoding json body"))
	}

	fmt.Println("Stored myRecipe")
	fmt.Println(myRecipe)
}
