package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	serviceLicense "github.com/isophtalic/License/internal/service/license"
)

func GetLicenses() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		data, page, totalPages := serviceLicense.GetLicenses(c.Request.URL.Query())
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

func GetLicenseByID() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == 0 {
			customError.Throw(http.StatusMethodNotAllowed, "Need to license_id")
			return
		}
		data, err := serviceLicense.GetLicenseByID(id)
		if err != nil {
			customError.Throw(http.StatusUnprocessableEntity, err.Error())
			return
		}

		Response(c, http.StatusOK, ResponseBody{
			Message: "successfully",
			Data:    data,
		})
	})
}

func CreateLicense() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		cmd := new(dto.CreateLicenseDTO)
		if err := c.BindJSON(&cmd); err != nil {
			customError.Throw(http.StatusUnprocessableEntity, err.Error())
			return
		}
		serviceLicense.CreateLicense(cmd)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Successfully",
			Data:    cmd.Name,
		})
	})
}

func UpdateLicense() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		licenseID := c.Param("id")
		if len(licenseID) == 0 {
			customError.Throw(http.StatusBadRequest, "no records to update")
			return
		}

		cmd := new(dto.UpdateLicenseDTO)
		if err := c.BindJSON(&cmd); err != nil {
			customError.Throw(http.StatusUnprocessableEntity, err.Error())
			return
		}

		serviceLicense.UpdateLicense(cmd, licenseID)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Successfully",
			Data:    cmd.Name,
		})
	})
}
