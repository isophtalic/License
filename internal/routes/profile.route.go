package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/app"
)

func profileRouter(parent *gin.RouterGroup) {
	router := parent.Group("/profile")
	router.GET("", app.Profile())
	router.PUT("/change-password", app.ChangePassword())
	router.PUT("", app.UpdateProfile())
}
