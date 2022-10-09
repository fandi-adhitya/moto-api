package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Errors.Errors() != nil {
			c.Next()
			c.JSON(http.StatusNotFound, gin.H{
				"code":   http.StatusNotFound,
				"status": http.StatusText(http.StatusNotFound),
			})
		}
	}
}
