package service

import (
	"errors"
	"net/http"
	"time"

	"git.cyradar.com/license-manager/backend/internal/dto"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"git.cyradar.com/license-manager/backend/internal/validators"
)

func UpdateLicenseConfig(cmd *dto.UpdateAndCreateLicenseConfigDTO, configID string) (code int, err error) {
	// validate
	validators.ValidateCreateLicenseConfigStruct(cmd)

	licenseConfig, err := persistence.LicenseConfig().FindByID(configID)
	if err != nil && err.Error() == "record not found" {
		return http.StatusNotFound, errors.New("Don't know which record to update")
	}
	if err != nil && err.Error() != "record not found" {
		return http.StatusInternalServerError, err
	}

	if licenseConfig.Name != &cmd.Name {
		_, err = persistence.LicenseConfig().FindByName(cmd.Name)
		if err == nil {
			return http.StatusMethodNotAllowed, errors.New("name is unique")
		}
	}

	newLicenseConfig := &models.LicenseConfig{
		ConfigID:  licenseConfig.ConfigID,
		Name:      &cmd.Name,
		Key:       &cmd.Key,
		Value:     &cmd.Value,
		Status:    &cmd.Status,
		CreatedAt: licenseConfig.CreatedAt,
		UpdatedAt: time.Now(),
		LicenseID: licenseConfig.LicenseID,
	}
	return 200, persistence.LicenseConfig().UpdateLicense(newLicenseConfig)

}
