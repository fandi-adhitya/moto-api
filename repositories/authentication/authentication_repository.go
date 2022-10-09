package authentication

import (
	"github.com/fandi-adhitya/moto-api.git/models"
	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	SignIn(tx *gorm.DB, user models.User) (models.User, error)
	SignUp(tx *gorm.DB, user models.User) (models.User, error)
}
