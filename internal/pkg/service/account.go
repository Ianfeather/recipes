package service

import (
	"database/sql"
)

// GetAccountID returns the account ID for a user
func GetAccountID(db *sql.DB, userID string) (int, error) {
	var accountID int
	accountQuery := `SELECT account_id from account_user WHERE user_id = ?;`
	if err := db.QueryRow(accountQuery, userID).Scan(&accountID); err != nil {
		// TODO: Return an error of type unknown user
		return 0, err
	}
	return accountID, nil
}
