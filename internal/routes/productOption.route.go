package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/isophtalic/License/internal/app"
)

func productOptionRouter(parent *gin.RouterGroup) {
	router := parent.Group("/product-option")
	router.POST("", app.CreateOptions())
	router.GET("/:id", app.DetailProductOption())
	router.PATCH("/:id", app.UpdateProductOption())
	router.DELETE("/:id", app.DeleteProductOption())

	router.POST("/:id/detail/create", app.AddOptionDetail())
}
