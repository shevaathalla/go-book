package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	mysql "gobook/pkg/db/mysql"
	author "gobook/src/author/injector"
	book "gobook/src/book/injector"
	oauth "gobook/src/oauth/injector"
	publisher "gobook/src/publisher/injector"
	register "gobook/src/register/injector"
	rental "gobook/src/rental/injector"
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
	author.InitializeService(db).Route(&r.RouterGroup)
	publisher.InitializeService(db).Route(&r.RouterGroup)
	book.InitializeService(db).Route(&r.RouterGroup)
	oauth.InitializeService(db).Route(&r.RouterGroup)
	rental.InitializeService(db).Route(&r.RouterGroup)
	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")

}
