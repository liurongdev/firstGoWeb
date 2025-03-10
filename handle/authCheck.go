package handle

import "github.com/gin-gonic/gin"

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.AbortWithStatusJSON(400, gin.H{"message": "Authorization header required"})
			return
		}
		c.Next()
	}
}
