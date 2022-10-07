package application

import (
	"database/sql"
	"github.com/fandi-adhitya/moto-api.git/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func NewDatabase() *gorm.DB {
	dsn := os.Getenv("DB")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.PanicError(err)

	db := sql.DB{}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return DB
}
