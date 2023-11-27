package infrastructure

import (
	"github.com/flexuxs/clubHubApp/src/domain"
)

type UserRepository struct {
	// Your database connection or any other necessary dependencies
}

func NewUserRepository() domain.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUser() *domain.User {
	// Retrieve user from the database
	return &domain.User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}
}
