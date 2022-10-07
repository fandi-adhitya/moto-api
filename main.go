package main

import (
	"github.com/fandi-adhitya/moto-api.git/application"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	application.NewDatabase()

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
