package routes

import (
	"database/sql"
	"supercotacao/backend/controllers"
	"supercotacao/backend/middleware"

	"github.com/gin-gonic/gin"
)

func Program(engine *gin.Engine, db *sql.DB) {
	api := engine.Group("v1")
	api.Use(middleware.ValidadeToken())
	api.POST("/novo-programa", func(context *gin.Context) {
		controllers.ControllerPostProgram(context, db)
	})
}
