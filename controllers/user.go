package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)


type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) ResisterUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("RegisterUser Called in UserController")
	uc.UserService.CreateUser()
	w.Write([]byte("User Registration Endpoint"))
}