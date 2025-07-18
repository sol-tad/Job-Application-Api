package handler

import (
	"database/sql"
	"jobApplication/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByIDHandler(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		id,err:=strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		user, err := services.GetUserByID(db, id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Unable to fetch user"})
			return
		}

		c.JSON(200, gin.H{"user": user})
	}

}
