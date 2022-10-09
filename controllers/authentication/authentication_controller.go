package authentication

import "github.com/gin-gonic/gin"

type AuthenticationController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}
