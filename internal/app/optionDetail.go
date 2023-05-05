package app

import (
	"net/http"

	optionDetailService "git.cyradar.com/license-manager/backend/internal/service/optionDetail"
	"github.com/gin-gonic/gin"
)

// [DELETE] /option_detail/:id
func DeleteOptionDetail() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		optionDetailID := c.Param("id")
		optionDetailService.DeleteByID(optionDetailID)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Delete option detail successfully",
		})
	})
}
