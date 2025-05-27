package controllers

import (
	"database/sql"
	"supercotacao/backend/models"
	"supercotacao/backend/services"

	"github.com/gin-gonic/gin"
)

func PostUser(context *gin.Context, db *sql.DB) {
	var data = models.User{}

	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(400, gin.H{
			"erro": "Erro ao ler informações para criar novo usuário. tente novamente",
			"code": 400,
		})
		return
	}

	data, err = services.CreateUser(data, db)
	if err != nil {
		context.JSON(400, gin.H{
			"erro": err.Error(),
			"code": 400,
		})
		return
	}

	context.JSON(201, gin.H{
		"user": data,
		"code": 201,
	})
}
