package validators

import (
	"net/http"

	"git.cyradar.com/license-manager/backend/internal/dto"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"github.com/asaskevich/govalidator"
)

func ValidateCustomer(customer *dto.UpdateAndAddCustomerDTO) {
	_, err := govalidator.ValidateStruct(customer)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ValidateEmail(customer.Email)
	ValidateName(customer.Name)
	ValidateName(customer.Company)
	ValidateOrganization(customer.Organization)
	ValidatePhoneNumber(customer.PhoneNumber)
}

func ValidateOrganization(organization string) {
	if organization == "personal" || organization == "enterprise" || organization == "government" {
		return
	}
	customError.Throw(http.StatusUnprocessableEntity, "organization is invalid")
}
