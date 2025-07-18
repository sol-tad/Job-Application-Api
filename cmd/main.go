package main

import (
	"jobApplication/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db,err:=repository.InitDB()

	if err!=nil{
		log.Fatal(err)
	}
	
	r := gin.Default()
	log.Println("rrr-->",r)
	r.Run(":8080")

}