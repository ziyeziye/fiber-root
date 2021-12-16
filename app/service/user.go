package service

import (
	"fiber-root/app/model"
	"fiber-root/db"
	"fiber-root/pkg/errno"
)

var UserService = &userService{}

type userService struct{}

func (u *userService) Login(user *model.User) error {
	result := db.GetDB().Where(user).First(user)
	if result.Error != nil {
		return errno.UserNotExits
	}
	return nil
}

func (u *userService) CheckAuth(username, password string) bool {
	//var auth Auth
	//db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	//if auth.ID > 0 {
	//	return true
	//}

	return false
}
