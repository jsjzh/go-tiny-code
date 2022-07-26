package controller

import (
	"go-tiny-code/src/model"
	"go-tiny-code/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createUser struct {
}

func CreateUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		if err := service.CreateUser(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			ctx.JSON(http.StatusCreated, user)
		}
	}
}

func DeleteUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		if err := service.DeleteUser(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}

func UpdateUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		if err := service.UpdateUser(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}

func ReadUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)

	user := model.User{
		Base: model.Base{
			ID: id,
		},
		Name:  ctx.Query("name"),
		Email: ctx.Query("Email"),
		Phone: ctx.Query("Phone"),
	}

	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		if err := service.ReadUser(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}

}
