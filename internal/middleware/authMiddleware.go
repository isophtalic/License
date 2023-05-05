package middleware

import (
	"context"
	"net/http"
	"strings"

	appController "git.cyradar.com/license-manager/backend/internal/app"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(JWTKEY string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var claims jwt.MapClaims
		header := ctx.GetHeader("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(header), "Bearer"))

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWTKEY), nil
		})

		if err != nil || token == nil || !token.Valid {
			ctx.Header("Content-Type", "Token invalid")
			appController.Response(ctx, http.StatusUnauthorized, appController.ResponseBody{
				Message: "Unauthorize",
			})
			ctx.Abort()
			return
		}

		keyRedis := claims["aud"].(string) + ":" + claims["jti"].(string)
		_, err = persistence.Account().Get(keyRedis)
		if err != nil {
			ctx.Header("Content-Type", "Token valid")
			appController.Response(ctx, http.StatusUnauthorized, appController.ResponseBody{
				Message: "Token is not suitable for this request session",
			})
			ctx.Abort()
			return
		}

		ctx.Request = ctx.Request.WithContext(context.WithValue(context.Background(), "user.email", claims["sub"].(string)))
		ctx.Next()
	}
}
