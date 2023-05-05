package app

import (
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// [GET] /api/v1/example/test
func ExampleController() func(*gin.Context) {
	return HandleErrorWrapper(func(c *gin.Context) {

		customError.Throw(http.StatusBadGateway, "Throw an error")

		Response(c, http.StatusOK, ResponseBody{
			Message: "OKELA",
		})
	})
}
