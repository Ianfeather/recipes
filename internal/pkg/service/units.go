package service

import (
	"database/sql"
)

// Unit is used to constrain ingredients
type Unit struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// GetAllUnits returns all unit types
func GetAllUnits(db *sql.DB) ([]Unit, error) {
	results, err := db.Query("SELECT id, name FROM unit order by lower(name);")

	if err != nil {
		return nil, err
	}

	units := make([]Unit, 0)

	for results.Next() {
		r := Unit{}
		err = results.Scan(&r.ID, &r.Name)
		if err != nil {
			return nil, err
		}
		units = append(units, r)
	}
	return units, nil
}
