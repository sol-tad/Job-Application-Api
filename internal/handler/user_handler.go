package handler

import (
	"database/sql"
	"jobApplication/internal/services"
	"net/http"
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

func UpdateUserProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var userUpdate struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}

		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")

		if !isAdmin && userID != id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to update this user profile"})
			return
		}

		updateUser, err := services.UpdateUserProfile(db, id, userUpdate.Username, userUpdate.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user profile"})
			return
		}

		c.JSON(http.StatusOK, updateUser)
	}
}