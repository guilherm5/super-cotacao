package main

import (
	"log"
	"supercotacao/backend/database"
	"supercotacao/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{
		"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD",
	}
	config.AllowHeaders = []string{
		"Origin", "Content-Type", "Authorization", "X-Requested-With", "Accept", "X-Access-Token",
	}
	gin.ForceConsoleColor()
	router := gin.Default()
	router.Use(cors.New(config))

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("erro: ", err)
		return
	}

	routes.User(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Println("falha ao iniciar o servidor: ", err)
	} else {
		log.Println("servidor rodando na porta 8080")
	}
}
