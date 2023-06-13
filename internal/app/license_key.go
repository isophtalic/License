package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	customError "github.com/isophtalic/License/internal/error"
	serviceLicenseKey "github.com/isophtalic/License/internal/service/license_key"
)

func MakeLicenseKey() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		id := c.Param("license_id")
		if len(id) == 0 {
			customError.Throw(http.StatusMethodNotAllowed, "Need to license_id")
			return
		}
		encodedCipherText, key := serviceLicenseKey.Encrypt(id)

		Response(c, http.StatusOK, ResponseBody{
			Message: "Successfully",
			Data: map[string]string{
				"Key":         key,
				"License_key": encodedCipherText,
			},
		})
	})
}

func GetLicenseKeys() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		data, page, totalPages := serviceLicenseKey.GetKeys(c.Request.URL.Query(), c)
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

func Active() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		err := serviceLicenseKey.Active(c)
		if err != nil {
			customError.Throw(http.StatusUnprocessableEntity, "Can't active license_key : "+err.Error())
		}
		Response(c, http.StatusOK, ResponseBody{
			Message: "successfully",
		})
	})
}
