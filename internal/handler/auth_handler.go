package handler

import (
	"database/sql"
	"jobApplication/internal/models"
	"jobApplication/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := services.LoginUser(db, user.Username, user.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
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

func ForgotPasswordHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.ForgotPasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		passowrd, err := services.ForgotPassword(db, req.Username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"updated_password": passowrd})
	}
}