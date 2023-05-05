package routes

import (
	app "git.cyradar.com/license-manager/backend/internal/app"
	"github.com/gin-gonic/gin"
)

func userRouter(parent *gin.RouterGroup) {
	router := parent.Group("/users")
	router.GET("", app.Users())
	router.PUT("/:id/change-password", app.ChangePassword())
	router.POST("/", app.AddNewUser())
	router.PUT("/:id", app.UpdateProfile())
}
