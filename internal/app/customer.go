package app

import (
	"net/http"

	"git.cyradar.com/license-manager/backend/internal/dto"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"github.com/gin-gonic/gin"

	serviceCustomer "git.cyradar.com/license-manager/backend/internal/service/customer"
)

func GetCustomers() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		data, page, totalPages := serviceCustomer.GetCustomers(c.Request.URL.Query())
		Response(c, http.StatusOK, ResponseBody{
			Message: "successfully",
			Data: map[string]interface{}{
				"page":       page,
				"totalPages": totalPages,
				"data":       data,
			},
		})
	})
}

func CreateCustomer() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		cmd := new(dto.UpdateAndAddCustomerDTO)
		if err := c.BindJSON(&cmd); err != nil {
			customError.Throw(http.StatusUnprocessableEntity, err.Error())
			return
		}

		err := serviceCustomer.AddCustomer(cmd)
		if err != nil {
			customError.Throw(http.StatusNotFound, err.Error())
			return
		}

		Response(c, http.StatusOK, ResponseBody{
			Message: "Successfully",
			Data:    cmd.Email,
		})

	})
}

func UpdateCustomer() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		customerID := c.Param("id")
		if len(customerID) == 0 {
			customError.Throw(http.StatusBadRequest, "no records to update")
			return
		}
		cmd := new(dto.UpdateAndAddCustomerDTO)
		if err := c.BindJSON(&cmd); err != nil {
			customError.Throw(http.StatusUnprocessableEntity, err.Error())
			return
		}

		err := serviceCustomer.UpdateCustomer(customerID, cmd)
		if err != nil {
			customError.Throw(http.StatusUnprocessableEntity, err.Error())
			return
		}
		Response(c, http.StatusOK, ResponseBody{
			Message: "Successfully",
			Data:    cmd.Email,
		})
	})
}
