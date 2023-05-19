package app

import (
	"net/http"
	"strings"

	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/helpers"
	serviceUser "github.com/isophtalic/License/internal/service/user"

	"github.com/gin-gonic/gin"
)

// POST /auth/sign-in
func Login() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		cmd := new(serviceUser.AuthSignIn)
		if err := c.BindJSON(&cmd); err != nil {
			Response(c, http.StatusUnprocessableEntity, ResponseBody{
				Message: err.Error(),
			})
			return
		}

		token, err := serviceUser.LogIn(cmd)
		if err != nil {
			Response(c, http.StatusBadRequest, ResponseBody{
				Message: "Login Unsuccessfully",
			})
			return
		}

		Response(c, http.StatusOK, ResponseBody{
			Message: "Login Successfully",
			Data: map[string]interface{}{
				"access_token": token,
			},
		})
	})

}

// POST /user/change-password
func ChangePassword() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {

		id := c.Param("id")

		data := new(dto.ChangePassword)
		if err := c.BindJSON(&data); err != nil {
			Response(c, http.StatusUnprocessableEntity, ResponseBody{
				Message: err.Error(),
			})
			return
		}

		header := c.GetHeader("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(header), "Bearer"))
		err := serviceUser.ChangePassword(data, id, tokenString)
		if err != nil {
			Response(c, http.StatusBadRequest, ResponseBody{
				Message: err.Error(),
			})
			return
		}

		Response(c, http.StatusOK, ResponseBody{
			Message: "Successfully",
		})
	})

	// redirect??
	// c.Redirect(http.StatusSeeOther, "/")
}

func SignUp() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		newUser := new(dto.AddUserDTO)
		if err := c.BindJSON(&newUser); err != nil {
			Response(c, http.StatusUnprocessableEntity, ResponseBody{
				Message: err.Error(),
			})
			return
		}

		header := c.GetHeader("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(header), "Bearer"))

		err := serviceUser.AddUser(tokenString, newUser)
		if err != nil {
			Response(c, http.StatusUnprocessableEntity, ResponseBody{
				Message: err.Error(),
			})
			return
		}

		Response(c, http.StatusOK, ResponseBody{
			Message: "Add user successfully",
		})
	})

}

func LogOut() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(header), "Bearer"))

		RedisKey := helpers.GetRedisKeyFromToken(tokenString)
		err := serviceUser.LogOut(RedisKey)
		if err != nil {
			Response(c, http.StatusUnprocessableEntity, ResponseBody{
				Message: err.Error(),
			})
			return
		}
		Response(c, http.StatusOK, ResponseBody{
			Message: "Log out successfully",
		})
	})
}
