package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	optionDetailService "github.com/isophtalic/License/internal/service/optionDetail"
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
