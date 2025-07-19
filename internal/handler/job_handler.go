package handler

import (
	"database/sql"
	"jobApplication/internal/models"
	"jobApplication/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateJobHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var job models.Job

		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		job.UserID = userID

		cratedJob, err := services.CreateJob(db, &job)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, cratedJob)
	}
}