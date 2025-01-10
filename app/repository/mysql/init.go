package mysql

import (
	"database/sql"
)

func InitRepositories(db *sql.DB) map[string]interface{} {
	repos := map[string]interface{}{
		"userRepo": NewUserRepository(db),
		// Add other repository instances here...
	}
	return repos
}
