package routes

import (
	"database/sql"
	"supercotacao/backend/controllers"

	"github.com/gin-gonic/gin"
)

func User(engine *gin.Engine, db *sql.DB) {
	engine.POST("/novo-usuario", func(context *gin.Context) {
		controllers.ControllerPostUser(context, db)
	})
}
