package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App will hold the dependencies of the application
type App struct{}

// NewApp returns the application itself
func NewApp() (*App, error) {
	app := &App{}
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
	router.HandleFunc("/recipe", a.recipeHandler)
	router.Use(loggingMiddleware)
	return router, nil
}
