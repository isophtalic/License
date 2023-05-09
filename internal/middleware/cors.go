package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/configs"
)

func CORS(config *configs.Configure) gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{config.ClientOrigin}
	return cors.New(cfg)
}
