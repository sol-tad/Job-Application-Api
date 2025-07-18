package handler

import (
	"database/sql"
	"jobApplication/internal/models"
	"jobApplication/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {
		var user models.User
		if err:=c.ShouldBindJSON(&user);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return 
		}
		
	}

}
func RegisterHandler(db *sql.DB) gin.HandlerFunc{

	return func(c *gin.Context) {
		var user models.User
		if err:=c.ShouldBindJSON(&user);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return 
		}
		err:=services.RegisterUser(db, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

	}

}

	