package validators

import (
	"net/http"
	"regexp"
	"time"

	"git.cyradar.com/license-manager/backend/internal/dto"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"git.cyradar.com/license-manager/backend/internal/helpers"
	"github.com/asaskevich/govalidator"
)

func ValidateCreateLicenseStruct(cmd *dto.CreateLicenseDTO) {
	_, err := govalidator.ValidateStruct(cmd)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ValidateName(cmd.Name)
	ValidateName(cmd.CustomerName)
	ValidateName(cmd.ProductName)
	validateOption(cmd.Option)
	validateTags(cmd.Tag)
}

func ValidateUpdateLicenseStruct(cmd *dto.UpdateLicenseDTO) {
	_, err := govalidator.ValidateStruct(cmd)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ValidateName(cmd.Name)
	validateOption(cmd.Option)
	validateTags(cmd.Tag)
}

func validateOption(option string) {
	if option == "Trial" || option == "Pro" {
		return
	}
	customError.Throw(http.StatusUnprocessableEntity, "option not exist : "+option)
}

func validateTags(tags []string) {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	for _, v := range tags {
		if !regex.MatchString(v) {
			customError.Throw(http.StatusUnprocessableEntity, "Name mustn't have special charater")
		}
	}
}

func ValidateExpireDate(date string) time.Time {
	if date == "" {
		return time.Now().Add(time.Hour * 8640)
	}

	dateParsed := helpers.ParseDate(date)
	if dateParsed.Day() == time.Now().Day() {
		customError.Throw(http.StatusMethodNotAllowed, "The expiration time is too short")
	}

	return dateParsed
}

func ValidateCreateLicenseConfigStruct(cmd *dto.UpdateAndCreateLicenseConfigDTO) {
	_, err := govalidator.ValidateStruct(cmd)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ValidateName(cmd.Name)
	validateKey(cmd.Key)
}

func validateKey(key string) {
	if key == "" {
		customError.Throw(http.StatusMethodNotAllowed, "Need to key")
	}
	if len(key) >= 254 {
		customError.Throw(http.StatusMethodNotAllowed, "Key is too long")
	}
}
