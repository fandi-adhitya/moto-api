package application

import (
	controller "github.com/fandi-adhitya/moto-api.git/controllers/authentication"
	"github.com/fandi-adhitya/moto-api.git/exceptions"
	repository "github.com/fandi-adhitya/moto-api.git/repositories/authentication"
	service "github.com/fandi-adhitya/moto-api.git/services/authentication"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, validate *validator.Validate) {
	router := gin.Default()

	authenticationRepository := repository.NewAuthenticationRepositoryImpl()
	authenticationService := service.NewAuthenticationServiceImpl(authenticationRepository, db, validate)
	authenticationController := controller.NewAuthenticationControllerImpl(authenticationService)

	router.Use(gin.Recovery())
	router.Use(exceptions.ErrorHandler())

	router.POST("/sign-in", authenticationController.SignIn)

	router.Run()
}
