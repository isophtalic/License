package validators

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(user *dto.AddUserDTO) {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		customError.Throw(http.StatusBadRequest, err.Error())
	}
	ValidateEmail(user.Email)
	ValidateStronglyPassword(user.Password)
	ValidateName(user.Name)
	ValidateRole(user.Role)
}

func ValidateUpdateProfile(cmd *dto.UpdateUserDTO) {
	_, err := govalidator.ValidateStruct(cmd)
	if err != nil {
		panic(err)
	}
	ValidateEmail(cmd.Email)
	ValidateName(cmd.Name)
	ValidateRole(cmd.Role)
}

func ValidateStronglyPassword(password string) error {
	// Kiểm tra độ dài mật khẩu
	if len(password) < 8 {
		customError.Throw(http.StatusUnprocessableEntity, "Minimum password length is 8")
	}

	// Kiểm tra chữ hoa
	re := regexp.MustCompile(`[A-Z]`)
	if !re.MatchString(password) {
		customError.Throw(http.StatusUnprocessableEntity, "Password must have a capital letter")
	}

	// Kiểm tra số
	re = regexp.MustCompile(`[0-9]`)
	if !re.MatchString(password) {
		customError.Throw(http.StatusUnprocessableEntity, "Password must have a number")
	}

	// Kiểm tra kí tự đặc biệt
	re = regexp.MustCompile(`[!@#$%^&*()]`)
	if !re.MatchString(password) {
		customError.Throw(http.StatusUnprocessableEntity, "Password must have a special charater")
	}

	return nil
}

func ValidateRole(role string) {
	if role != "admin" {
		role = "customer"
	}
	if role == "admin" || role == "customer" {
		return
	}
	customError.Throw(http.StatusUnprocessableEntity, "Invalid role")
}

func ValidateName(name string) {
	if len(name) < 3 {
		customError.Throw(http.StatusUnprocessableEntity, "Minimum of name is 3")
	}

	regex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !regex.MatchString(name) {
		customError.Throw(http.StatusUnprocessableEntity, "Name mustn't have special charater")
	}
}
func ValidateFormChangePassword(data *dto.ChangePassword) error {

	if data.Password != "" {
		if err := ValidateStronglyPassword(data.Password); err != nil {
			return err
		}

	}
	if err := ValidateStronglyPassword(data.NewPassword); err != nil {
		return err
	}

	return nil
}

func VerifyPassword(oldPassword, newPassword, salt, HashedPassword string) error {

	if newPassword == oldPassword {
		return errors.New("The new password cannot be the same as the old password")
	}
	OldPassByte := append([]byte(oldPassword), []byte(salt)...)
	if err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), OldPassByte); err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}

func IsDesc(order string) bool {
	if order == "DESC" || order == "desc" {
		return true
	}
	return false
}
