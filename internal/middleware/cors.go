package middleware

import (
	"git.cyradar.com/license-manager/backend/internal/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(config *configs.Configure) gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{config.CLIENT_ORIGIN}
	return cors.New(cfg)
}
