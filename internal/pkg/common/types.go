package common

import "database/sql"

// Env is passed into our application
type Env struct {
	DB *sql.DB
}
