package services

import (
	"database/sql"
	"jobApplication/internal/models"
	"jobApplication/internal/repository"
)

func GetUserByID(db *sql.DB, id int) (*models.User, error) {

	return repository.GetUserByID(db, id)
}