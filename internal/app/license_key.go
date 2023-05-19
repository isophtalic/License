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
