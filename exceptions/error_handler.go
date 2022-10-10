package exceptions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(c *gin.Context, error interface{}) {
	if notFoundError(c, error) {
		return
	}

	if emailAlreadyExist(c, error) {
		return
	}

	if validationError(c, error) {
		return
	}

	internalServerError(c, error)
}

func validationError(c *gin.Context, e interface{}) bool {
	if exception, ok := e.(validator.ValidationErrors); ok {

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
			"error":   exception.Error(),
		})

		return true
	} else {
		return false
	}
}

func emailAlreadyExist(c *gin.Context, e interface{}) bool {
	if exception, ok := e.(EmailAlreadyExist); ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
			"error":   exception.Error,
		})

		return true
	} else {
		return false
	}
}

func notFoundError(c *gin.Context, error interface{}) bool {
	exception, ok := error.(NotFoundError)
	fmt.Println(ok)
	if ok {

		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": http.StatusText(http.StatusNotFound),
			"error":   exception.Error,
		})

		return true
	} else {
		return false
	}
}

func internalServerError(c *gin.Context, e interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": http.StatusText(http.StatusInternalServerError),
		"error":   e,
	})
	return
}
