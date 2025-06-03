package routes

import (
	"database/sql"
	"supercotacao/backend/controllers"

	"github.com/gin-gonic/gin"
)

func Login(engine *gin.Engine, db *sql.DB) {
	engine.POST("/login-usuario", func(context *gin.Context) {
		controllers.GenerateToken(context, db)
	})
}
