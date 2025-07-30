package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID() (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db:_db,
	}
}

func (u *UserRepositoryImpl) GetByID() (*models.User, error) {
	fmt.Println("Fetching User in UserRepository")

	//step 1: prepare the query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"

	//step 2: Execute the query
	row := u.db.QueryRow(query, 1)

	// step 3: Process the row data and create an output object (Process the result)
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Println("Error Scanning user:", err)
			return nil, err
		}
	} 

	//step 4: Print the user details
	fmt.Println("User fetched Successfully:", user)

	return user, nil
}
