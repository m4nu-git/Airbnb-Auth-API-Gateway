package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID() (*models.User, error)
	Create(username string, email string, hashedPassword string) error
	GetByEmail(email string) (*models.User, error)
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

func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, username, email, hashedPassword)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User created succesfully, rows affected:", rowsAffected)
	return nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = ?"

	row := u.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("Error Scanning user:", err)
			return nil, err
		}
	}
	return user, nil
}
