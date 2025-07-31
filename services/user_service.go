package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching User in UserService")
	u.userRepository.GetByID()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	password := "example_password"

	hashedPassword, err := utils.HashedPassword(password)
	if err != nil {
		return err
	}

	u.userRepository.Create(
		"username_example1",
		"user@example1.com",
		hashedPassword,
	)
	return nil
}

func (u *UserServiceImpl) LoginUser() (string, error) {
	// Pre - requisites: This function will be given email and password as parameter, which we can hardcode for now.
	// response := utils.CheckPasswordHash("example_password", "$2a$10$ozMKQPgKbCW7vx41zYjAKOyeSTG9RPhO6aLhgZPQmP/BtpXUF.YOS")
	// fmt.Println("Login response:", response)

	email := "user@example1.com"
	password := "example_password"
	
	// Step 1. make a repository call to get the user by email
	user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("Error fetching user by email")
		return "", err
	}

	// Step 2. If user exists, or not. If nor exists, return error
	if user == nil {
		fmt.Println("No user found with the given email")
		return "", fmt.Errorf("No user foundwith email: %s", email)
	}

	// Step 3. If user exists, check the password using utils.CheckPasswordHash

	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", nil
	}

	// Step 4. If Password matches, print a JWT token, else return error

	payload := jwt.MapClaims{
		"email": user.Email,
		"id": user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("JWT Token:", tokenString)
 
	return tokenString, nil
}