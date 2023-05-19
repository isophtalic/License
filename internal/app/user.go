package app

import (
	"net/http"
	"strings"

	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/models"
	serviceUser "github.com/isophtalic/License/internal/service/user"

	"github.com/gin-gonic/gin"
)

// GET /api/v1/users
func Users() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(header), "Bearer"))

		data, page, totalPages := serviceUser.GetUsers(tokenString, c.Request.URL.Query())
		dataResponse := make([]models.ResponseUser, 0)
		for _, v := range data {
			dataResponse = append(dataResponse, *v.ToResponseUser())
		}
		Response(c, http.StatusOK, ResponseBody{
			Message: "successfully",
			Data: map[string]interface{}{
				"page":       page,
				"totalPages": totalPages,
				"data":       dataResponse,
			},
		})
	})
}

// POST api/v1/users/add
func AddNewUser() gin.HandlerFunc {
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

// POST /api/v1/users/update
func UpdateProfile() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		id := c.Param("id")

		data := new(dto.UpdateUserDTO)
		if err := c.BindJSON(&data); err != nil {
			Response(c, http.StatusUnprocessableEntity, ResponseBody{
				Message: err.Error(),
			})
			return
		}

		mail := c.Request.Context().Value("user.email").(string)
		err := serviceUser.UpdateProfile(data, mail, id)
		if err != nil {
			Response(c, http.StatusBadRequest, ResponseBody{
				Message: err.Error(),
			})
			return
		}

		Response(c, http.StatusOK, ResponseBody{
			Message: "Update succesfully",
		})

		return
	})
}

func Profile() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		email := c.Request.Context().Value("user.email").(string)
		user := serviceUser.Profile(email)
		userResponse := user.ToResponseUser()
		Response(c, http.StatusOK, ResponseBody{
			Message: "successfully",
			Data: map[string]interface{}{
				"data": &userResponse,
			},
		})
	})
}
