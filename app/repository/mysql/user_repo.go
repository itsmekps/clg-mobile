package mysql

import (
	"database/sql"

	"fiber-boilerplate/app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec(`
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?);
	`, user.Username, user.Email, user.Password)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(`
		SELECT id, username, email, password
		FROM users
		WHERE username = ?;
	`, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(id int) (*models.User, error) {
	var user = models.User{}
	err := r.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
