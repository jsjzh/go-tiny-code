package service

import (
	"go-tiny-code/src/model"
)

func CreateUser(user *model.User) error {
	if err := model.DB().Create(&user).Error; err != nil {
		return err
	} else {
		model.DB().Take(&user)
	}
	return nil
}

func DeleteUser(user *model.User) error {
	if err := model.DB().Delete(&user).Error; err != nil {
		return err
	} else {
		model.DB().Take(&user)
	}
	return nil
}

func UpdateUser(user *model.User) error {
	if err := model.DB().Model(&user).Updates(&user).Error; err != nil {
		return err
	} else {
		model.DB().Take(&user)
	}
	return nil
}

func ReadUser(user *model.User) error {
	if err := model.DB().Take(&user).Error; err != nil {
		return err
	}

	return nil
}
