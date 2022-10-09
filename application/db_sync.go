package application

import "github.com/fandi-adhitya/moto-api.git/models"

func SyncDB() {
	NewDB().AutoMigrate(&models.User{})
}
