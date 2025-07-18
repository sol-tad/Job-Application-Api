package routes

import (
	"database/sql"
	"jobApplication/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine,db *sql.DB){
	//Auth Routes
	r.POST("/login",handler.LoginHandler(db))
	r.POST("/register",handler.RegisterHandler(db))
	r.GET("/users/:id",handler.GetUserByIDHandler(db))



}