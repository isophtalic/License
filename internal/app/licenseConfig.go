package app

import (
	"net/http"

	"git.cyradar.com/license-manager/backend/internal/dto"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	serviceLicenseConfig "git.cyradar.com/license-manager/backend/internal/service/licenseConfig"
	"github.com/gin-gonic/gin"
)

func GetLicenseConfigs() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		data, page, totalPages := serviceLicenseConfig.GetLicenseConfigs(c)
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

func CreateConfig() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		licenseID := c.Param("license_id")
		if licenseID == "" {
			customError.Throw(http.StatusMethodNotAllowed, "Need to license_id")
		}

		cmd := new(dto.UpdateAndCreateLicenseConfigDTO)
		if err := c.BindJSON(&cmd); err != nil {
			customError.Throw(http.StatusUnprocessableEntity, err.Error())
			return
		}

		serviceLicenseConfig.CreateLicenseConfig(cmd, licenseID)

		Response(c, http.StatusOK, ResponseBody{
			Message: "Successfully",
			Data:    cmd.Name,
		})
	})
}

func UpdateConfig() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		configID := c.Param("id")
		if configID == "" {
			customError.Throw(http.StatusMethodNotAllowed, "Need to license_id")
		}
		cmd := new(dto.UpdateAndCreateLicenseConfigDTO)
		if err := c.BindJSON(&cmd); err != nil {
			customError.Throw(http.StatusUnprocessableEntity, err.Error())
			return
		}

		code, err := serviceLicenseConfig.UpdateLicenseConfig(cmd, configID)
		if err != nil {
			customError.Throw(code, err.Error())
			return
		}

		Response(c, http.StatusOK, ResponseBody{
			Message: "Successfully",
			Data:    cmd.Name,
		})
	})
}
