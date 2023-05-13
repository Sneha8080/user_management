// user_service.go (services package)

package services

import (
	"github.com/sneha8080/userManagement/models"
	repositories "github.com/sneha8080/userManagement/repository"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) CreateUser(user *models.User) error {
	// Additional validation or business logic can be added here before creating the user.
	return us.userRepository.CreateUser(user)
}

func (us *UserService) GetUserByID(id int) (*models.User, error) {
	return us.userRepository.GetUserByID(id)
}
