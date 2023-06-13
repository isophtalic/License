package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/app"
)

func customerRouter(parent *gin.RouterGroup) {
	router := parent.Group("/customers")
	router.GET("", app.GetCustomers())
	router.POST("", app.CreateCustomer())
	router.PUT("/:id", app.UpdateCustomer())
}
