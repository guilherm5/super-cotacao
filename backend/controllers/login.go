package controllers

import (
	"database/sql"
	"supercotacao/backend/models"
	"supercotacao/backend/utils"

	"supercotacao/backend/services"

	"github.com/gin-gonic/gin"
)

func GenerateToken(context *gin.Context, db *sql.DB) {
	var dataLogin models.User

	err := context.ShouldBindJSON(&dataLogin)
	if err != nil {
		utils.Logger.Error("Erro ao ler body para iniciar login ", err)
		context.JSON(400, gin.H{
			"erro": "dados inválidos",
			"code": 400,
		})
		return
	}

	token, err := services.CompareAndGenerateToken(dataLogin, db)
	if err != nil {
		context.JSON(401, gin.H{
			"erro": err.Error(),
			"code": 400,
		})
		return
	}

	context.JSON(200, gin.H{
		"login": "Bearer " + token,
		"code":  200,
	})
}
