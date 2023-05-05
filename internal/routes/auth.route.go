package routes

import (
	app "git.cyradar.com/license-manager/backend/internal/app"
	"github.com/gin-gonic/gin"
)

func authRouter(parent *gin.RouterGroup) {
	router := parent.Group("/auth")
	router.POST("/log-in", app.Login())
	router.POST("/log-out", app.LogOut())
}
