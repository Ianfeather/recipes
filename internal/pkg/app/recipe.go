package app

import (
	"net/http"
)

func (a *App) recipeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Recipe page"))
}
