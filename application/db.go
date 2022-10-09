package application

import (
	"github.com/fandi-adhitya/moto-api.git/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func NewDB() *gorm.DB {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.PanicError(err)

	//db := &sql.DB{}
	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)
	//db.SetConnMaxIdleTime(5 * time.Minute)
	//db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
