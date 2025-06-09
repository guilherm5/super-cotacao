package routes

import (
	"database/sql"
	"supercotacao/backend/controllers"
	"supercotacao/backend/middleware"

	"github.com/gin-gonic/gin"
)

func Tree(engine *gin.Engine, db *sql.DB) {
	api := engine.Group("v1")
	api.Use(middleware.ValidadeToken())
	api.POST("/nova-arvore", func(context *gin.Context) {
		controllers.ControllerPostTree(context, db)
	})
}
