package app

import (
	"encoding/json"
	"net/http"
)

// Recipe contains recipe fields
type Recipe struct {
	Name        string
	Ingredients int
}

func (a *App) recipeHandler(w http.ResponseWriter, req *http.Request) {
	recipe := Recipe{"French Toast", 4}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(recipe)
	if err != nil {
		w.Write([]byte("Error encoding json"))
	}
}
