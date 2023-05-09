package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/isophtalic/License/internal/app"
)

func authRouter(parent *gin.RouterGroup) {
	router := parent.Group("/auth")
	router.POST("/log-in", app.Login())
	router.POST("/log-out", app.LogOut())
}
