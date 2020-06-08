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
func (a *App) GetRouter() (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/health", healthHandler).Methods("GET")
	router.HandleFunc("/recipe/{slug}", a.recipeHandler).Methods("GET")
	router.HandleFunc("/recipe", a.addRecipeHandler).Methods("POST")
	router.HandleFunc("/shopping-list", a.getListHandler).Methods("GET")
	router.Use(loggingMiddleware)
	return router, nil
}
