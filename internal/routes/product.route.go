package routes

import (
	app "git.cyradar.com/license-manager/backend/internal/app"
	"git.cyradar.com/license-manager/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func productRouter(parent *gin.RouterGroup) {
	router := parent.Group("/product")
	// product
	router.POST("/", app.CreateProduct())
	router.GET("/", app.ListProducts())
	router.PATCH("/:id", app.UpdateProduct())
	router.GET("/:id", app.DetailProduct())
	//adjkf
	router.GET("/search/:keyword", app.Search())
	router.PATCH("/:id/status", app.ChangeProductStatus())

	//key
	router.POST("/:id/key", app.GenerateNewKeys())
	router.POST("/:id/key/upload",
		middleware.SizeLimiterMiddleware(5000),
		middleware.ExtRestrictMiddleware([]string{".pem"}),
		app.UploadKeys(),
	)
}
