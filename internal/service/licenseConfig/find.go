package service

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
)

func GetLicenseConfigs(c *gin.Context) (licenseConfigs []models.LicenseConfig, page int, total_pages int) {
	perPge := strings.TrimSpace(c.Query("per_page"))
	pge := strings.TrimSpace(c.Query("page"))
	sort := strings.TrimSpace(c.Query("sort"))
	search := strings.TrimSpace(c.Query("search"))
	licenseID := c.Param("license_id")
	var err error

	if licenseID == "" {
		customError.Throw(http.StatusMethodNotAllowed, "Need to license_id")
	}

	// search
	if search != "" {
		licenseConfigs, page, total_pages, err = persistence.LicenseConfig().Search(search, licenseID, perPge, pge, sort)
		if err != nil {
			customError.Throw(http.StatusNotFound, errors.New("Not Found").Error())
			return
		}
		return
	}

	// select all
	licenseConfigs, page, total_pages, err = persistence.LicenseConfig().GetConfigsByLicenseId(perPge, pge, sort, licenseID)
	if err != nil {
		customError.Throw(http.StatusNotFound, errors.New("Not Found").Error())
		return
	}
	return
}
