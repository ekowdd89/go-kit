package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)


func ZeroNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		start:= time.Now()
		c.Next()
		latency:= time.Since(start)
		statusCode:= c.Writer.Status()
		log.Info().Str("client_ip", c.ClientIP()).Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", statusCode).
			Str("latency", latency.String()).
			Str("user_agent", c.Request.UserAgent()).
			Msg("incoming request")
	}
}