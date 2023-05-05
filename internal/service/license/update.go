package service

import (
	"net/http"
	"strings"
	"time"

	"git.cyradar.com/license-manager/backend/internal/dto"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"git.cyradar.com/license-manager/backend/internal/validators"
	"github.com/lib/pq"
)

func UpdateLicense(cmd *dto.UpdateLicenseDTO, licenseID string) {
	validators.ValidateUpdateLicenseStruct(cmd)

	license, err := persistence.License().FindById(licenseID)
	if err != nil {
		customError.Throw(http.StatusFound, "Invalid license")
		return
	}

	if license.Name != &cmd.Name {
		_, err = persistence.License().FindByName(cmd.Name)
		if err != nil {
			customError.Throw(http.StatusFound, "Name was used")
			return
		}
	}

	tags, _ := pq.Array(cmd.Tag).Value()
	expireDate := validators.ValidateExpireDate(cmd.ExpiryDate)
	newLicense := &models.License{
		LicenseID:   license.LicenseID,
		Name:        &cmd.Name,
		Option:      &cmd.Option,
		Description: &cmd.Description,
		Tag:         strings.Fields(tags.(string)),
		ExpiryDate:  expireDate,
		CreatedAt:   license.CreatedAt,
		UpdatedAt:   time.Now(),
		CustomerID:  license.CustomerID,
		ProductID:   license.ProductID,
	}
	err = persistence.License().UpdateLicense(newLicense)
	if err != nil {
		customError.Throw(http.StatusInternalServerError, err.Error())
	}
}
