package service

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/isophtalic/License/internal/validators"
)

func CreateLicenseConfig(cmd *dto.UpdateAndCreateLicenseConfigDTO, licenseID string) {
	_, err := persistence.License().FindById(licenseID)
	if err != nil && err.Error() == "record not found" {
		customError.Throw(http.StatusMethodNotAllowed, "license is invalid")
		return
	}
	if err != nil && err.Error() != "record not found" {
		customError.Throw(http.StatusMethodNotAllowed, err.Error())
		return
	}

	_, err = persistence.LicenseConfig().FindByName(cmd.Name)
	if err == nil {
		customError.Throw(http.StatusMethodNotAllowed, "name is unique")
		return
	}

	id := uuid.NewString()
	validators.ValidateCreateLicenseConfigStruct(cmd)
	newLicenseConfig := &models.LicenseConfig{
		ConfigID:  &id,
		Name:      &cmd.Name,
		Key:       &cmd.Key,
		Value:     &cmd.Value,
		Status:    &cmd.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LicenseID: &licenseID,
	}
	persistence.LicenseConfig().Save(newLicenseConfig)
}
