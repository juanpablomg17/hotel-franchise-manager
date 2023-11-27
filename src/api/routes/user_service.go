package routes

import (
	"net/http"

	"github.com/flexuxs/clubHubApp/src/domain"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUser(c *gin.Context) {
	// Call UserRepository to retrieve user

	// add middleware to check
	user := s.userRepository.GetUser()

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
