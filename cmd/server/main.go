package main

import (
	"jobApplication/internal/repository"
	"jobApplication/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}


	db,err:=repository.InitDB()
    defer db.Close()
	if err!=nil{
		log.Fatal(err)
	}
	
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	
	r := gin.Default()
	routes.InitRoutes(r, db)
	log.Println("rrr-->",r)
	r.Run(":"+port)

}