package service

import (
	"fmt"

	"database/sql"
)

// Ingredient is a subset of shopping List
type Ingredient struct {
	Unit      string  `json:"unit"`
	Quantity  float64 `json:"quantity"`
	IsChecked bool    `json:"isChecked"`
}

// ListItem is used to interface with the DB
type ListItem struct {
	Name      string
	Unit      string
	Quantity  float64
	IsChecked bool
}

func removeIngredientListItems(userID int, db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM list WHERE user_id = ? AND type = 'ingredient';")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID)
	if err != nil {
		fmt.Println("could not delete ingredients")
		return err
	}
	return nil
}

func addIngredientListItems(userID int, ingredients map[string]*Ingredient, db *sql.DB) error {
	sqlStr := "INSERT INTO list(user_id, name, type, quantity, is_checked, unit_id) VALUES "
	vals := []interface{}{}

	for name, val := range ingredients {
		sqlStr += "(?, ?, 'ingredient', ?, false, (SELECT id from unit where name='?')),"
		vals = append(vals, userID, name, val.Quantity, val.Unit)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(vals...)
	if err != nil {
		fmt.Println("could not add ingredients to shopping list")
		return err
	}
	return nil
}

func getExtraListItems(userID int, db *sql.DB) (map[string]*Ingredient, error) {
	ingredientQuery := "SELECT name, unit.name as unit, quantity, is_checked as isChecked FROM list INNER JOIN unit on unit_id = unit.id WHERE user_id = ? and type = 'extra';"
	results, err := db.Query(ingredientQuery, userID)

	items := make([]ListItem, 0)

	for results.Next() {
		item := ListItem{}
		err = results.Scan(&item.Name, &item.Unit, &item.Quantity, &item.IsChecked)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	extrasList := make(map[string]*Ingredient)
	for _, item := range items {
		newItem := Ingredient{
			Unit:      item.Unit,
			Quantity:  item.Quantity,
			IsChecked: item.IsChecked,
		}
		extrasList[item.Name] = &newItem
	}

	return extrasList, nil
}
