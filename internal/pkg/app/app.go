package app

import (
	"database/sql"
	"log"
	"net/http"
	"recipes/internal/pkg/common"

	"github.com/gorilla/mux"
)

// App will hold the dependencies of the application
type App struct {
	db *sql.DB
}

// NewApp returns the application itself
func NewApp(env *common.Env) (*App, error) {
	app := &App{
		db: env.DB,
	}
	return app, nil
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ok"))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.RequestURI)
		next.ServeHTTP(w, req)
	})
}

// GetRouter returns the application router
func (a *App) GetRouter(base string) (*mux.Router, error) {
	router := mux.NewRouter()
	router.HandleFunc(base+"/health", healthHandler).Methods("GET")
	router.HandleFunc(base+"/recipes", a.recipesHandler).Methods("GET")
	router.HandleFunc(base+"/ingredients", a.ingredientsHandler).Methods("GET")
	router.HandleFunc(base+"/recipe/{slug:[a-zA-Z-]+}", a.recipeHandlerBySlug).Methods("GET")
	router.HandleFunc(base+"/recipe/{id:[0-9]+}", a.recipeHandlerByID).Methods("GET")
	router.HandleFunc(base+"/recipe", a.addRecipeHandler).Methods("POST")
	router.HandleFunc(base+"/recipe", a.editRecipeHandler).Methods("PUT")
	router.HandleFunc(base+"/shopping-list", a.getListHandler).Methods("GET")
	router.HandleFunc(base+"/shopping-list", a.createListHandler).Methods("POST")
	router.HandleFunc(base+"/shopping-list/buy", a.buyListItemHandler).Methods("PATCH")
	router.HandleFunc(base+"/shopping-list/extra", a.addExtraListItem).Methods("POST")
	router.HandleFunc(base+"/shopping-list/clear", a.clearListHandler).Methods("DELETE")
	router.HandleFunc(base+"/units", a.getUnitsHandler).Methods("GET")
	router.Use(loggingMiddleware)
	return router, nil
}
