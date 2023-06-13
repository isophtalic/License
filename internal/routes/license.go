package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/app"
)

func licenseRouter(parent *gin.RouterGroup) {
	router := parent.Group("/license")
	router.GET("", app.GetLicenses())
	router.POST("", app.CreateLicense())
	router.PUT("/:id", app.UpdateLicense())
	// api handle config
	router.GET("/:license_id/config", app.GetLicenseConfigs())
	router.POST("/:license_id/config", app.CreateConfig())
	router.PUT("/config/:id", app.UpdateConfig())

	// api make license key
	router.GET("/:license_id/license_key", app.GetLicenseKeys())
	router.POST("/:license_id/license_key", app.MakeLicenseKey())
	router.POST("/:license_id/license_key/:license_key", app.Active())
}
