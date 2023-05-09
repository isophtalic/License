package validators

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
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
