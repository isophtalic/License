package routes

import (
	"git.cyradar.com/license-manager/backend/internal/app"
	"github.com/gin-gonic/gin"
)

func customerRouter(parent *gin.RouterGroup) {
	router := parent.Group("/customers")
	router.GET("", app.GetCustomers())
	router.POST("/", app.CreateCustomer())
	router.PUT("/:id", app.UpdateCustomer())
}
