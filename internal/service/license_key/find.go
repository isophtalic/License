package service

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"

	customError "github.com/isophtalic/License/internal/error"
)

func GetKeys(query url.Values, c *gin.Context) (licenseKeys []models.License_key, page, total_pages int) {
	perPge := strings.TrimSpace(query.Get("per_page"))
	pge := strings.TrimSpace(query.Get("page"))
	sort := strings.TrimSpace(query.Get("sort"))

	LicenseID := c.Param("license_id")
	if len(LicenseID) == 0 {
		customError.Throw(http.StatusMethodNotAllowed, "Need to license_id")
		return
	}

	licenseKeys, page, total_pages, err := persistence.LicenseKey().GetLicenseKey(perPge, pge, sort, LicenseID)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, "Can't get license-key from database")
	}
	return
}
