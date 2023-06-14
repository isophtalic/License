package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/isophtalic/License/internal/app"
	"github.com/isophtalic/License/internal/middleware"
)

func productRouter(parent *gin.RouterGroup) {
	router := parent.Group("/product")
	// product
	router.POST("", app.CreateProduct())
	router.GET("", app.ListProducts())
	router.PATCH("/:id", app.UpdateProduct())
	router.GET("/:id", app.DetailProduct())
	//adjkf
	router.GET("/search/:keyword", app.Search())
	router.PATCH("/:id/status", app.ChangeProductStatus())

	//key
	router.POST("/:id/key", app.GenerateNewKeys())
	router.GET("/:id/key", app.GetKeyProduct())
	router.POST("/:id/key/upload",
		middleware.SizeLimiterMiddleware(5000),
		middleware.ExtRestrictMiddleware([]string{".pem"}),
		app.UploadKeys(),
	)
}
