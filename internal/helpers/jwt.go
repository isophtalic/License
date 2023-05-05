package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"git.cyradar.com/license-manager/backend/internal/configs"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	JwtKey string
}

var expirationTime = time.Now().Add(time.Hour * 24).Unix()

type Claims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

func (JS JWTService) IssueJWT(data interface{}) (string, error) {
	claims := Claims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime,
		},
	}
	// token, := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return CreateToken(claims, JS.JwtKey)
}

func (JS JWTService) ValidateJWT(tokenString string, data interface{}) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(JS.JwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New("invalid signature")
		}
	}
	if !token.Valid {
		return errors.New("token not valid")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return JS.parseDataFromMap(claims, data)
	}

	return fmt.Errorf("invalid JWTService")
}

func (JS JWTService) parseDataFromMap(m jwt.MapClaims, out interface{}) error {
	data, ok := m["data"]
	if !ok || data == nil {
		return fmt.Errorf("invalid JWTService Claims: Data")
	}

	buffers, e := json.Marshal(data)
	if e != nil {
		return e
	}

	return json.Unmarshal(buffers, out)
}
func CreateToken(claims jwt.Claims, JwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JwtSecret))
}
func NewJWT(Key string) *JWTService {
	service := new(JWTService)
	service.JwtKey = Key
	return service
}

func decode(tokenString string) (jwt.MapClaims, error) {
	config, err := configs.GetConfig()
	if err != nil {
		return nil, err
	}
	// check existence of user
	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET_KEY), nil
	})

	if err != nil || tkn == nil || !tkn.Valid {
		return nil, err
	}
	return claims, nil
}

func DecodeToken(tokenString string) (email string, role string, err error) {
	claims, err := decode(tokenString)
	if err != nil {
		return "", "", err
	}
	return claims["sub"].(string), claims["role"].(string), nil
}

func GetRedisKeyFromToken(token string) (RedisKey string) {
	claims, err := decode(token)
	if err != nil {
		customError.Throw(http.StatusMethodNotAllowed, err.Error())
	}
	keyRedis := claims["aud"].(string) + ":" + claims["jti"].(string)
	return keyRedis
}
