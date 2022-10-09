package authentication

import (
	"github.com/fandi-adhitya/moto-api.git/models/web"
)

type AuthenticationService interface {
	SignIn(request web.AuthRequest) (web.AuthResponse, error)
	SignUp(request web.AuthRequest) (web.AuthResponse, error)
}
