package routes

import (
	"git.cyradar.com/license-manager/backend/internal/app"
	"github.com/gin-gonic/gin"
)

func licenseRouter(parent *gin.RouterGroup) {
	router := parent.Group("/license")
	router.GET("", app.GetLicenses())
	router.POST("/", app.CreateLicense())
	router.PUT("/:id", app.UpdateLicense())
	// api handle config
	router.GET("/:license_id/config", app.GetLicenseConfigs())
	router.POST("/:license_id/config", app.CreateConfig())
	router.PUT("/config/:id", app.UpdateConfig())
}
