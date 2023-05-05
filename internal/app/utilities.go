package app

import (
	"fmt"
	"net/http"
	"reflect"

	customError "git.cyradar.com/license-manager/backend/internal/error"
	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, code int, data ResponseBody) {
	c.JSON(code, gin.H{
		"Message": data.Message,
		"Data":    data.Data,
	})
}

/*
HandleErrorWrapper wrap controller and response an existing error.
*/
func HandleErrorWrapper(controller func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}

			fmt.Println("error:", err)
			switch reflect.TypeOf(err) {
			case reflect.TypeOf(customError.Type{}):
				Response(c, customError.Cast(err).StatusCode, ResponseBody{
					Message: customError.Cast(err).Error(),
				})
				return
			default:
				Response(c, http.StatusInternalServerError, ResponseBody{
					Message: "Something went wrong",
				})
			}
		}()

		controller(c)
	}

}
