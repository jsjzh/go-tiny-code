package controller

import (
	"go-tiny-code/src/model"
	"go-tiny-code/src/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(ctx *gin.Context) {
	service.CreateUser()
}

func DeleteUser(ctx *gin.Context) {
	service.DeleteUser()
}

func UpdateUser(ctx *gin.Context) {
	service.UpdateUser()
}

func ReadUser(ctx *gin.Context) {

	query := model.User{
		Name: ctx.Query("name"),
	}

	if err := validator.New().StructCtx(ctx, query); err != nil {
		log.Println(err.Error())
	}

	// if result, err := service.ReadUser(query); err != nil {
	// 	log.Println(err)
	// }

	ctx.JSON(200, query)

}
