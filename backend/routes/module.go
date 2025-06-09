package routes

import (
	"database/sql"
	"supercotacao/backend/controllers"
	"supercotacao/backend/middleware"

	"github.com/gin-gonic/gin"
)

func Module(engine *gin.Engine, db *sql.DB) {
	api := engine.Group("v1")
	api.Use(middleware.ValidadeToken())
	api.POST("/novo-modulo", func(context *gin.Context) {
		controllers.ControllerPostModule(context, db)
	})
}
