package authentication

import (
	"github.com/fandi-adhitya/moto-api.git/helpers"
	"github.com/fandi-adhitya/moto-api.git/models"
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

	token, err := controller.AuthenticationService.GenerateToken(auth)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

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

	token, err := controller.AuthenticationService.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   user,
	})

	return
}

func (controller *AuthenticationControllerImpl) Validate(c *gin.Context) {
	user, _ := c.Get("user")

	toResponseModel := user.(models.User)

	c.JSON(http.StatusOK, gin.H{
		"status": "Logged in",
		"user":   toResponseModel,
	})
}
