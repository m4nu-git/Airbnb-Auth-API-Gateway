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

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("GetUserById Called in UserController")
	uc.UserService.GetUserById()
	w.Write([]byte("User fetching Endpoint done"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("CreateUser Called in UserController")
	uc.UserService.CreateUser()
	w.Write([]byte("User fetching Endpoint done"))
}


func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("LoginUser Called in UserController")
	uc.UserService.LoginUser()
	w.Write([]byte("User login Endpoint done"))
}
