package service

import (
	"go-tiny-code/src/model"
	"log"
)

func CreateUser() {

	user := model.User{
		Name:  "king",
		Email: "kimimi_king@163.com",
		Phone: "18368094601",
	}

	if error := model.DB().Create(&user).Error; error != nil {
		log.Println(error)
	}

}

func DeleteUser() {

}

func UpdateUser() {

}

func ReadUser(query model.User) {
	user := map[string]interface{}{}

	model.DB().Take(&user)

	log.Println("user", user)

	// var user model.User

	// model.DB().Take(query)

	// return , nil
}
