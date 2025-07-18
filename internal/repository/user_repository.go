package repository

import (
	"database/sql"
	"jobApplication/internal/models"
)

func CreateUser(db *sql.DB, user *models.User) error {

	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	_, err := db.Exec(query, user.Username, user.Password, user.Email)
	
	return err
}