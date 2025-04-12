package middleware

import "github.com/gin-gonic/gin"

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.GetHeader("x-api-key")
		if apikey == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "API key is required",
			})
			return
		}
		if apikey != "jgk4jg325g325gj2k4hj2h4bjhhkhj234" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Invalid API key",
			})
		}
		c.Next()
	}
}
