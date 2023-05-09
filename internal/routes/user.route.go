package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/isophtalic/License/internal/app"
)

func userRouter(parent *gin.RouterGroup) {
	router := parent.Group("/users")
	router.GET("", app.Users())
	router.PUT("/:id/change-password", app.ChangePassword())
	router.POST("/", app.AddNewUser())
	router.PUT("/:id", app.UpdateProfile())
}
