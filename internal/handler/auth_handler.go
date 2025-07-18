package handler

import (
	"database/sql"
	"jobApplication/internal/models"
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

	}

}