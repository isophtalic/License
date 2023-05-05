package routes

import (
	app "git.cyradar.com/license-manager/backend/internal/app"
	"github.com/gin-gonic/gin"
)

func optionDetailRouter(parent *gin.RouterGroup) {
	router := parent.Group("/option-detail")
	router.DELETE("/:id", app.DeleteOptionDetail())
}
