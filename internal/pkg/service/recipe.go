package service

import (
	"recipes/internal/pkg/common"

	"database/sql"
	"fmt"
)

func getIngredientsByRecipeID(id int, db *sql.DB) ([]common.Ingredient, error) {
	ingredientQuery := `
		SELECT
			ingredient.name as name,
			unit.name as unit,
			quantity,
			department.name as department
		FROM
			part
			INNER JOIN ingredient on ingredient_id = ingredient.id
			INNER JOIN unit on unit_id = unit.id
			LEFT JOIN department on department.id = (select department_id from ingredient_department where ingredient_department.ingredient_id = ingredient.id)
		WHERE
		recipe_id = ?;
	`
	results, err := db.Query(ingredientQuery, id)

	ingredients := make([]common.Ingredient, 0)

	for results.Next() {
		var department sql.NullString
		ingredient := common.Ingredient{}
		err = results.Scan(&ingredient.Name, &ingredient.Unit, &ingredient.Quantity, &department)

		if err != nil {
			return nil, err
		}

		if department.Valid {
			ingredient.Department = department.String
		} else {
			ingredient.Department = ""
		}

		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

// GetRecipeBySlug fetches a recipe from the database by Slug
func GetRecipeBySlug(slug string, userID int, db *sql.DB) (r *common.Recipe, e error) {
	recipe := &common.Recipe{Ingredients: []common.Ingredient{}}
	recipeQuery := `
		SELECT recipe.id as id, name, remote_url
			FROM recipe_user
			INNER JOIN recipe on recipe_user.recipe_id = recipe.id
			WHERE recipe_user.slug= ? AND recipe_user.user_id = ?;`

	var remoteURL sql.NullString
	err := db.QueryRow(recipeQuery, slug, userID).Scan(&recipe.ID, &recipe.Name, &remoteURL)

	if err != nil {
		return nil, err
	}

	if remoteURL.Valid {
		recipe.RemoteURL = remoteURL.String
	}

	ingredients, err := getIngredientsByRecipeID(recipe.ID, db)

	if err != nil {
		return nil, err
	}

	recipe.Ingredients = ingredients

	return recipe, nil
}

// GetRecipeByID fetches a recipe from the database by ID
func GetRecipeByID(id int, userID int, db *sql.DB) (r *common.Recipe, e error) {
	recipe := &common.Recipe{Ingredients: []common.Ingredient{}}
	recipeQuery := `
		SELECT recipe.id as id, name, remote_url
			FROM recipe_user
			INNER JOIN recipe on recipe_user.recipe_id = recipe.id
			WHERE recipe_user.recipe_id= ? AND recipe_user.user_id = ?;`

	var remoteURL sql.NullString
	err := db.QueryRow(recipeQuery, id, userID).Scan(&recipe.ID, &recipe.Name, &remoteURL)

	if err != nil {
		return nil, err
	}

	if remoteURL.Valid {
		recipe.RemoteURL = remoteURL.String
	}

	ingredients, err := getIngredientsByRecipeID(id, db)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	recipe.Ingredients = ingredients
	return recipe, nil
}

// AddRecipe inserts recipe, ingredients into the DB
func AddRecipe(recipe common.Recipe, userID int, db *sql.DB) error {

	stmt, err := db.Prepare("INSERT INTO recipe (name, slug, remote_url) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(recipe.Name, common.Slugify(recipe.Name), recipe.RemoteURL)

	if err != nil {
		fmt.Println("could not insert recipe")
		return err
	}

	id, err := res.LastInsertId()
	recipe.ID = int(id)

	if err = insertIngredients(recipe, db); err != nil {
		return err
	}
	if err = insertParts(recipe, db); err != nil {
		return err
	}
	if err = addRecipeToUser(recipe, userID, db); err != nil {
		return err
	}
	return nil
}

// EditRecipe updates recipe information
func EditRecipe(recipe common.Recipe, userID int, db *sql.DB) error {
	var id string
	// Checking to see if this recipe exists for this user
	if err := db.QueryRow("SELECT id FROM recipe_user WHERE recipe_id=? AND user_id = ?;", recipe.ID, userID).Scan(&id); err == sql.ErrNoRows {
		fmt.Println("no results")
		return err
	} else if err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE recipe SET name=?, remote_url=? WHERE id=?")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(recipe.Name, recipe.RemoteURL, recipe.ID); err != nil {
		return err
	}

	if err = insertIngredients(recipe, db); err != nil {
		return err
	}

	// Delete the existing relationships between recipe & ingredients
	stmt, err = db.Prepare("DELETE FROM part WHERE recipe_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(recipe.ID)
	if err != nil {
		return err
	}

	if err = insertParts(recipe, db); err != nil {
		return err
	}
	return nil
}

// DeleteRecipe removes a recipe from the db
func DeleteRecipe(recipe common.Recipe, userID int, db *sql.DB) error {
	var id string
	// Checking to see if this recipe exists for this user
	if err := db.QueryRow("SELECT id FROM recipe_user WHERE recipe_id=? AND user_id = ?;", recipe.ID, userID).Scan(&id); err == sql.ErrNoRows {
		fmt.Println("no results")
		return err
	} else if err != nil {
		return err
	}
	// Delete the existing relationships between recipe & ingredients
	stmt, err := db.Prepare("DELETE FROM part WHERE recipe_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(recipe.ID)
	if err != nil {
		return err
	}

	// Delete the existing relationships between recipe & user
	stmt, err = db.Prepare("DELETE FROM recipe_user WHERE recipe_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(recipe.ID)
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("DELETE FROM recipe WHERE id=?")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(recipe.ID); err != nil {
		return err
	}

	return nil
}

func insertIngredients(recipe common.Recipe, db *sql.DB) error {
	ingredientQuery := "INSERT INTO ingredient (name) values "
	for idx, ingredient := range recipe.Ingredients {
		ingredientQuery += fmt.Sprintf("('%s')", ingredient.Name)
		if idx != len(recipe.Ingredients)-1 {
			ingredientQuery += ","
		}
	}
	ingredientQuery += " ON DUPLICATE KEY UPDATE id=id;"

	if _, err := db.Exec(ingredientQuery); err != nil {
		fmt.Println("could not insert ingredients")
		return err
	}
	return nil
}

func insertParts(recipe common.Recipe, db *sql.DB) error {
	partsQuery := "INSERT INTO part (recipe_id, ingredient_id, unit_id, quantity) VALUES "
	for idx, ingredient := range recipe.Ingredients {
		partsQuery += fmt.Sprintf("(%d, ", recipe.ID)
		partsQuery += fmt.Sprintf("(SELECT id FROM ingredient WHERE name = '%s'),", ingredient.Name)
		partsQuery += fmt.Sprintf("(SELECT id FROM unit WHERE name = '%s'),", ingredient.Unit)
		partsQuery += fmt.Sprintf("%s) ", ingredient.Quantity)
		if idx != len(recipe.Ingredients)-1 {
			partsQuery += ","
		} else {
			partsQuery += ";"
		}
	}

	_, err := db.Exec(partsQuery)
	if err != nil {
		fmt.Println("could not insert part")
		return err
	}

	return nil
}

func addRecipeToUser(recipe common.Recipe, userID int, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO recipe_user (recipe_id, recipe_slug, user_id) values (?,?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(recipe.ID, common.Slugify(recipe.Name), userID)
	if err != nil {
		fmt.Println("could not relate recipe to user")
		return err
	}
	return nil
}
