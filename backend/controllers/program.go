package controllers

import (
	"database/sql"
	"supercotacao/backend/models"
	"supercotacao/backend/services"

	"supercotacao/backend/utils"

	"github.com/gin-gonic/gin"
)

func ControllerPostProgram(context *gin.Context, db *sql.DB) {
	var data = models.Program{}

	err := context.ShouldBindJSON(&data)
	if err != nil {
		utils.Logger.Info("Erro ao validar body da requisição ", err)
		context.JSON(400, gin.H{
			"erro": "Erro ao ler informações para criar novo programa. tente novamente",
			"code": 400,
		})
		return
	}

	data.CreatedBy, err = utils.GetUserFromContextPtr(context)
	if err != nil {
		context.JSON(400, gin.H{
			"erro": err.Error(),
			"code": 400,
		})
		return
	}

	data, err = services.CreateProgram(data, db)
	if err != nil {
		utils.Logger.Info("Erro ao chamar service para criação de um novo programa ", err)
		context.JSON(400, gin.H{
			"erro": err.Error(),
			"code": 400,
		})
		return
	}

	context.JSON(201, gin.H{
		"program": data,
		"code":    201,
	})
}
