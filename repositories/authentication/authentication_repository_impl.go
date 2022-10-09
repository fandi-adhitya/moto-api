package authentication

import (
	"errors"
	"github.com/fandi-adhitya/moto-api.git/models"
	"gorm.io/gorm"
)

type AuthenticationRepositoryImpl struct {
}

func NewAuthenticationRepositoryImpl() *AuthenticationRepositoryImpl {
	return &AuthenticationRepositoryImpl{}
}

func (auth *AuthenticationRepositoryImpl) SignIn(tx *gorm.DB, user models.User) (models.User, error) {
	find := tx.First(&user, "email = ? ", user.Email)
	if find.Error != nil {
		return user, errors.New("Cannot find user")
	}
	return user, nil

}

func (auth *AuthenticationRepositoryImpl) SignUp(tx *gorm.DB, user models.User) (models.User, error) {
	panic("any")
}
