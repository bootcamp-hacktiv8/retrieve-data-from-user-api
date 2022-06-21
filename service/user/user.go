package user

import (
	_entities "latihan-golang-usamah/entities"
	_userRepository "latihan-golang-usamah/repository/user"
)

type UserService struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepository _userRepository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) GetAllUser() ([]_entities.User, error) {
	users, err := us.userRepository.GetAllUser()
	return users, err
}
