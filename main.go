package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	mysql "gobook/pkg/db/mysql"
)

func init() {
	pathdir, _ := os.Getwd()
	environment := godotenv.Load(filepath.Join(pathdir, ".env"))

	if environment != nil {
		panic(environment)
	}
}

func main() {
	r := gin.Default()

	mysql.DB() // <- yg ini

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// user.InitializeUserHandler(db).Route(&r.RouterGroup)
	r.Run("127.0.0.1:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
