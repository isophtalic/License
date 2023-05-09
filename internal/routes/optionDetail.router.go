package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/isophtalic/License/internal/app"
)

func optionDetailRouter(parent *gin.RouterGroup) {
	router := parent.Group("/option-detail")
	router.DELETE("/:id", app.DeleteOptionDetail())
}
