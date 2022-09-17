package repository

import (
	"fmt"
	"tweeter/db"

	"github.com/google/uuid"
)

var UserRepository = userRepository{}

type userRepository struct{}

func (u userRepository) CreateUser(username, email, password string) error {
	const CreateUserQuery = "INSERT INTO users (`id`, `email`, `username`, `password`)VALUES ( ?, ?, ?, ?);"
	id := uuid.New().String()
	result, err := db.DB.Exec(CreateUserQuery, id, email, username, password)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
