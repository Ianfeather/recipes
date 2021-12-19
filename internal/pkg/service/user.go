package service

import (
	"database/sql"
	"log"
	"recipes/internal/pkg/common"
)

func AddUser(db *sql.DB, user common.User) error {
	userQuery := `
		INSERT INTO user (id, name, email)
			VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE
				id=id,
				name=?,
				email=?,
				last_logged_in_at=CURRENT_TIMESTAMP
			;
	`
	_, err := db.Exec(userQuery, user.ID, user.Name, user.Email, user.Name, user.Email)
	if err != nil {
		log.Println("Error adding user")
		log.Println(err)
		return err
	}
	return nil
}
