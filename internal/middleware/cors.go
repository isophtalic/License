package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/configs"
)

func CORS(config *configs.Configure) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", config.ClientOrigin)
		// c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, PATCH, OPTIONS, GET, PUT")
		// c.Header("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
