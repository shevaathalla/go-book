package user

import (
	"gobook/src/middleware"
	userDto "gobook/src/user/dto"
	userService "gobook/src/user/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) Route(r *gin.RouterGroup) {
	userRouter := r.Group("/api/user")
	userRouter.Use(middleware.AuthJwt)
	{
		userRouter.GET("/", userHandler.FindAll)
		userRouter.GET("/:id", userHandler.FindById)
		userRouter.POST("/", userHandler.Create)
	}
}

func (userHandler *UserHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if offset < 0 || offset == 0 {
		offset = 1
	}

	if limit < 0 {
		limit = 10
	}

	data := userHandler.userService.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success fetch data all user",
		"data":    data,
	})
}

func (userHandler *UserHandler) FindById(ctx *gin.Context) {

	id, paramErr := strconv.Atoi(ctx.Param("id"))

	if paramErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id not valid",
			"data":    nil,
		})

		ctx.Abort()
		return
	}

	data, err := userHandler.userService.FindById(id)

	if err != nil {

		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "user not found",
				"data":    nil,
			})

			ctx.Abort()
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "there is something wrong",
				"data":    nil,
			})

			ctx.Abort()
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success fetch data user with id " + ctx.Param("id"),
		"data":    data,
	})
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {

	var input userDto.CreateUserRequest

	inputErr := ctx.ShouldBindJSON(&input)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputErr.Error(),
			"data":    nil,
		})

		ctx.Abort()
		return
	}

	data, err := userHandler.userService.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})

		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success create user",
		"data":    data,
	})

}
