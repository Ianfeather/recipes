package common

import "database/sql"

// Env is passed into our application
type Env struct {
	DB *sql.DB
}

// SimpleResponse only returns a status message
type SimpleResponse struct {
	Status string `json:"status"`
}

// Ingredient contains ingredient fields
type Ingredient struct {
	Name       string `json:"name"`
	Unit       string `json:"unit"`
	Quantity   string `json:"quantity"`
	Department string `json:"department"`
}

// Recipe contains recipe fields
type Recipe struct {
	Name        string       `json:"name"`
	ID          int          `json:"id"`
	RemoteURL   string       `json:"remoteUrl"`
	Ingredients []Ingredient `json:"ingredients"`
}

// ListIngredient is a subset of shopping List
type ListIngredient struct {
	Unit       string  `json:"unit"`
	Quantity   float64 `json:"quantity"`
	IsBought   bool    `json:"isBought"`
	RecipeID   int     `json:"recipe_id"`
	Department string  `json:"department"`
}

// User object
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

// Account holds accounts and users
type Account struct {
	ID    int    `json:"id"`
	Users []User `json:"users`
}
