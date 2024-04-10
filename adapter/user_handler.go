package http

import (
	"minnnano-schedule/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(su usecase.IUserUsecase) *userHandler {
	return &userHandler{
		usecase: su,
	}
}

func (sh *userHandler) FindUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c
		users, err := sh.usecase.FindUsers(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, users)
	}
}

func (sh *userHandler) FindUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		user, err := sh.usecase.FindUserById(ctx, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, user)
	}
}
