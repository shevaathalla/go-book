package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	mysql "gobook/pkg/db/mysql"
	register "gobook/src/register/injector"
	user "gobook/src/user/injector"
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

	db := mysql.DB() // <- yg ini

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	user.InitializeService(db).Route(&r.RouterGroup)
	register.InitializeService(db).Route(&r.RouterGroup)
	r.Run("127.0.0.1:8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")

}
