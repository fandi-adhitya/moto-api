package authentication

import (
	"github.com/fandi-adhitya/moto-api.git/models/web"
	"github.com/fandi-adhitya/moto-api.git/services/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationControllerImpl struct {
	AuthenticationService authentication.AuthenticationService
}

func NewAuthenticationControllerImpl(authenticationService authentication.AuthenticationService) *AuthenticationControllerImpl {
	return &AuthenticationControllerImpl{AuthenticationService: authenticationService}
}

func (controller *AuthenticationControllerImpl) SignIn(c *gin.Context) {
	payload := web.AuthRequest{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "validation error",
			"message": err.Error(),
		})
		return
	}

	auth, err := controller.AuthenticationService.SignIn(payload)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   http.StatusNotFound,
			"status": http.StatusText(http.StatusNotFound),
			"data":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   auth,
	})

}

func (controller *AuthenticationControllerImpl) SignUp(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
