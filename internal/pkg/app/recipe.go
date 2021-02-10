package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"recipes/internal/pkg/common"
	"recipes/internal/pkg/service"
	"strconv"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

func (a *App) recipeHandlerBySlug(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	slug := mux.Vars(req)["slug"]

	recipe, err := service.GetRecipeBySlug(slug, userID, a.db)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Recipe not found", http.StatusNotFound)
			return
		}
		fmt.Println(err)
		http.Error(w, "Failed to parse recipe from db", http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(recipe)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}
}

func (a *App) recipeHandlerByID(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	id, err := strconv.Atoi(mux.Vars(req)["id"])

	if err == nil {
		fmt.Println(err)
	}

	recipe, err := service.GetRecipeByID(id, userID, a.db)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Recipe not found", http.StatusNotFound)
			return
		}
		fmt.Println(err)
		http.Error(w, "Failed to parse recipe from db", http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(recipe)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}
}

func (a *App) addRecipeHandler(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	recipe := common.Recipe{}
	err := json.NewDecoder(req.Body).Decode(&recipe)

	if err != nil {
		w.Write([]byte("Error decoding json body"))
	}

	err = service.AddRecipe(recipe, userID, a.db)

	if err != nil {
		fmt.Println("could not insert ingredients")
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	err = encoder.Encode("ok")

	fmt.Printf("Stored %s", recipe.Name)
}

func (a *App) editRecipeHandler(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	recipe := common.Recipe{}
	err := json.NewDecoder(req.Body).Decode(&recipe)
	encoder := json.NewEncoder(w)

	if err != nil {
		w.Write([]byte("Error decoding json body"))
	}

	if recipe.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: missing id"))
	}

	err = service.EditRecipe(recipe, userID, a.db)

	if err != nil {
		fmt.Println("could not update recipe")
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	err = encoder.Encode("ok")
	fmt.Printf("Updated %s", recipe.Name)
}

func (a *App) deleteRecipeHandler(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	recipe := common.Recipe{}
	err := json.NewDecoder(req.Body).Decode(&recipe)
	encoder := json.NewEncoder(w)

	if err != nil {
		w.Write([]byte("Error decoding json body"))
	}

	if recipe.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: missing id"))
	}

	err = service.DeleteRecipe(recipe, userID, a.db)

	if err != nil {
		fmt.Println("could not delete recipe")
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	err = encoder.Encode("ok")
	fmt.Printf("Updated %s", recipe.Name)
}
