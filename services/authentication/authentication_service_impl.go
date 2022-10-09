package authentication

import (
	"github.com/fandi-adhitya/moto-api.git/helpers"
	"github.com/fandi-adhitya/moto-api.git/models"
	"github.com/fandi-adhitya/moto-api.git/models/web"
	"github.com/fandi-adhitya/moto-api.git/repositories/authentication"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthenticationServiceImpl struct {
	AuthenticationRepository authentication.AuthenticationRepository
	DB                       *gorm.DB
	Validate                 *validator.Validate
}

func NewAuthenticationServiceImpl(authenticationRepository authentication.AuthenticationRepository, DB *gorm.DB, validate *validator.Validate) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{AuthenticationRepository: authenticationRepository, DB: DB, Validate: validate}
}

func (service *AuthenticationServiceImpl) SignIn(request web.AuthRequest) (web.AuthResponse, error) {
	err := service.Validate.Struct(request)
	helpers.PanicError(err)

	tx := service.DB.Begin()
	helpers.PanicError(tx.Error)
	defer helpers.CommitOrRollback(tx)

	auth := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	response, err := service.AuthenticationRepository.SignIn(tx, auth)

	if err != nil {
		return web.AuthResponse{}, err
	}

	return web.AuthResponse{
		Id:        response.ID,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}, nil
}

func (service *AuthenticationServiceImpl) SignUp(request web.AuthRequest) (web.AuthResponse, error) {
	//TODO implement me
	panic("implement me")
}
