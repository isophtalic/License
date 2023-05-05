package routes

import (
	app "git.cyradar.com/license-manager/backend/internal/app"
	"github.com/gin-gonic/gin"
)

func productOptionRouter(parent *gin.RouterGroup) {
	router := parent.Group("/product-option")
	router.POST("/", app.CreateOptions())
	router.GET("/:id", app.DetailProductOption())
	router.PATCH("/:id", app.UpdateProductOption())
	router.DELETE("/:id", app.DeleteProductOption())

	router.POST("/:id/detail/create", app.AddOptionDetail())
}
