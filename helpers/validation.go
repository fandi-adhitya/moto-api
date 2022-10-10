package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidationJson(c *gin.Context, payload interface{}) {
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "validation error",
			"message": err.Error(),
		})
		return
	}
}
