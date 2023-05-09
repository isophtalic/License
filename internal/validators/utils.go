package validators

import (
	"net/http"
	"net/mail"
	"regexp"

	"github.com/google/uuid"
	customError "github.com/isophtalic/License/internal/error"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, err.Error())
	}

	re := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	ok := re.MatchString(email)
	if !ok {
		customError.Throw(http.StatusUnprocessableEntity, "Email is invalid")
	}

	return nil
}

func ValidateUUID(u string) {
	_, err := uuid.Parse(u)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, err.Error())
	}
}

func ValidatePhoneNumber(phoneNumber string) {
	re := regexp.MustCompile(`^[0-9]{10,}$`)
	ok := re.MatchString(phoneNumber)
	if !ok {
		customError.Throw(http.StatusUnprocessableEntity, "Phone number is invalid")
	}
}
