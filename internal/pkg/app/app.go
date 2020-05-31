package app

import (
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

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Success from the handler"))
}

// GetRouter returns the application router
func (a *App) GetRouter() (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler).Methods("GET")
	return router, nil
}
