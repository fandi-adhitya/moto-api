package main

import (
	"github.com/fandi-adhitya/moto-api.git/application"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	db := application.NewDB()
	validate := validator.New()
	application.SyncDB()

	application.NewRouter(db, validate)

}
