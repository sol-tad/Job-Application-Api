package routes

import (
	"database/sql"
	"jobApplication/internal/auth"
	"jobApplication/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine,db *sql.DB){
	//Auth Routes
	r.POST("/login",handler.LoginHandler(db))
	r.POST("/register",handler.RegisterHandler(db))

	// User routes // employer
	authenticated := r.Group("/")
	authenticated.Use(auth.AuthMiddleware())
	authenticated.GET("/users/:id",handler.GetUserByIDHandler(db))
	authenticated.PUT("/users/:id", handler.UpdateUserProfileHandler(db))
	authenticated.POST("/users/:id/picture", handler.UpdateUserProfilePictureHandler(db))



}