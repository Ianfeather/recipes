package app

import (
	"recipes/internal/pkg/service"

	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/form3tech-oss/jwt-go"
)

func (a *App) recipesHandler(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	encoder := json.NewEncoder(w)
	recipes, err := service.GetAllRecipes(a.db, userID)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Recipe not found", http.StatusNotFound)
			err = encoder.Encode(make([]string, 0))
			return
		}
		fmt.Println(err)
		http.Error(w, "Failed to get recipes from db", http.StatusInternalServerError)
		return
	}

	err = encoder.Encode(recipes)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}
}
