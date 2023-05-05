package service

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.cyradar.com/license-manager/backend/internal/configs"
	"git.cyradar.com/license-manager/backend/internal/helpers"
	providerJWT "git.cyradar.com/license-manager/backend/internal/helpers"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"git.cyradar.com/license-manager/backend/internal/validators"

	customError "git.cyradar.com/license-manager/backend/internal/error"
	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const redisKey = "user_access"

var numberOfFailLogin = 0

type MyClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Name string `json:"names"`
}

type AuthSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLogOut struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func (account *AuthSignIn) Valid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}

	validators.ValidateEmail(account.Email)
	if err := validators.ValidateStronglyPassword(account.Password); err != nil {
		return err
	}

	return nil
}

func LogIn(cmd *AuthSignIn) (string, error) {
	// find user in db
	var now = time.Now()

	err := cmd.Valid()
	if err != nil {
		return "", err
	}

	config, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}

	user, err := persistence.User().SearchByEmail(cmd.Email)
	if err != nil {
		return "", errors.New("Account not exist")
	}

	if !(*user.Status) {
		return "", errors.New("Account has been disabled")
	}

	if err := validateNumberOfFailLogin(); err != nil {
		customError.Throw(http.StatusConflict, err.Error())
	}

	if !helpers.CompareAndHashPassword(*user.Password, cmd.Password, *user.Salt) {
		persistence.Account().Increment("numberOfFailLogin" + *user.Email)
		return "", errors.New("Password is not correct")
	}

	claims := cmd.makeClaims(*user, now)

	tokenString, err := providerJWT.CreateToken(claims, config.JWT_SECRET_KEY)
	if err != nil {
		return "", errors.New("Cannot create token")
	}

	keyRedis := *user.UserID + ":" + claims.Id
	value := strings.Split(tokenString, ".")[2]
	err = persistence.Account().Set(keyRedis, value, time.Now().Add(24*time.Hour))
	if err != nil {
		customError.Throw(http.StatusConflict, "Unable to create token")
	}

	setLoginSuccess(user)

	return tokenString, nil
}

func LogOut(key string) error {
	return persistence.Account().DeleteAll(key)
}

/*
Make a claims
*/
func (cmd *AuthSignIn) makeClaims(user models.User, c time.Time) MyClaims {

	return MyClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			Audience:  *user.UserID,
			Subject:   *user.Email,
			IssuedAt:  c.Unix(),
			ExpiresAt: c.Add(time.Hour * 24).Unix(),
		},
		Role: *user.Role,
		Name: *user.Name,
	}
}

/*
Set the number of failed login attempts to 0 and update the last login
*/
func setLoginSuccess(user *models.User) {
	err := persistence.Account().Delete("numberOfFailLogin" + *user.Email)
	if err != nil {
		customError.Throw(http.StatusConflict, err.Error())
	}

	// update last login
	err = persistence.User().UpdateLastLogin(*user)
	if err != nil {
		customError.Throw(http.StatusConflict, err.Error())
	}

}

/*
Check the number of failures,
return error if you get error more than 5 times
*/
func validateNumberOfFailLogin() error {
	numberFail, _ := persistence.Account().Get("numberOfFailLogin")
	if numberFail == "" {
		numberFail = string(rune(0))
	}
	numberFailInt, _ := strconv.Atoi(numberFail)
	if numberFailInt > 5 {
		return errors.New("You logged in wrong more than 5 times. Try again after 5 minute")
	}
	return nil
}
