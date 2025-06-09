package controllers

import (
	"database/sql"
	"supercotacao/backend/models"
	"supercotacao/backend/services"

	"supercotacao/backend/utils"

	"github.com/gin-gonic/gin"
)

func ControllerPostUser(context *gin.Context, db *sql.DB) {
	var data = models.User{}

	err := context.ShouldBindJSON(&data)
	if err != nil {
		utils.Logger.Error("Erro ao validar body da requisição ", err)
		context.JSON(400, gin.H{
			"erro": "Erro ao ler informações para criar novo usuário. tente novamente",
			"code": 400,
		})
		return
	}

	data, err = services.CreateUser(data, db)
	if err != nil {
		utils.Logger.Error("Erro ao chamar service para criação do usuário ", err)
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
