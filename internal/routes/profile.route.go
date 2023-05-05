package routes

import (
	"git.cyradar.com/license-manager/backend/internal/app"
	"github.com/gin-gonic/gin"
)

func profileRouter(parent *gin.RouterGroup) {
	router := parent.Group("/profile")
	router.GET("/", app.Profile())
	router.PUT("/change-password", app.ChangePassword())
	router.PUT("", app.UpdateProfile())
}
