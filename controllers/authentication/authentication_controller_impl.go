package authentication

import (
	"github.com/fandi-adhitya/moto-api.git/helpers"
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

	helpers.ValidationJson(c, &payload)

	auth := controller.AuthenticationService.SignIn(payload)

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   auth,
	})

	return

}

func (controller *AuthenticationControllerImpl) SignUp(c *gin.Context) {
	payload := web.AuthRequest{}

	helpers.ValidationJson(c, &payload)

	user := controller.AuthenticationService.SignUp(payload)

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   user,
	})

	return
}
