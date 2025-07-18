package services

import (
	"database/sql"
	"jobApplication/internal/models"
	"jobApplication/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB,user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
   
    if err != nil {
        return err
    }


    user.Password = string(hashedPassword)
    return repository.CreateUser(db, user)

    
}