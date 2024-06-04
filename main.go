package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pedropmartiniano/mvc-golang/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()

	routes.InitRoutes(&server.RouterGroup)

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
