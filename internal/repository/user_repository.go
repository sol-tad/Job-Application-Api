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

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	var profilePicture sql.NullString // Use sql.NullString to handle NULL values
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &profilePicture)
	if err != nil {
		return nil, err
	}
	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	} else {
		user.ProfilePicture = nil
	}
	return &user, nil
}