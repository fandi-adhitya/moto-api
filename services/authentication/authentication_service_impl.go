package authentication

import (
	"github.com/fandi-adhitya/moto-api.git/exceptions"
	"github.com/fandi-adhitya/moto-api.git/helpers"
	"github.com/fandi-adhitya/moto-api.git/models"
	"github.com/fandi-adhitya/moto-api.git/models/web"
	"github.com/fandi-adhitya/moto-api.git/repositories/authentication"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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

func (service *AuthenticationServiceImpl) SignIn(request web.AuthRequest) web.AuthResponse {
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
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return web.AuthResponse{
		Id:        response.ID,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
}

func (service *AuthenticationServiceImpl) SignUp(request web.AuthRequest) web.AuthResponse {
	err := service.Validate.Struct(request)
	helpers.PanicError(err)

	tx := service.DB.Begin()
	helpers.PanicError(tx.Error)
	defer helpers.CommitOrRollback(tx)

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helpers.PanicError(err)

	user := models.User{
		Email:    request.Email,
		Password: string(password),
	}

	response, err := service.AuthenticationRepository.SignUp(tx, user)

	if err != nil {
		panic(exceptions.NewEmailAlreadyExist(err.Error()))
	}

	return web.AuthResponse{
		Id:        response.ID,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
}
