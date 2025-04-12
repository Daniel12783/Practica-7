package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Antes del handler
		startTime := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIp := c.ClientIP()
		log.Printf("Request: %s %s from %s", method, path, clientIp)
		c.Next()
		//Despues del handler
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		statusCode := c.Writer.Status()
		log.Printf("Response: %d %s in %v", statusCode, path, duration)
	}
}
