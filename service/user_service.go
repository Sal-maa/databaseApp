package service

import (
	"designpattern/entity"
	"designpattern/repository"
	"fmt"
)

type UserService interface {
	LoginUserService(login entity.Login) (entity.User, error)
	GetUsersService() ([]entity.User, error)
	GetUserByIdService(id int) (entity.User, error)
	GetUserByNameService(name string) (entity.User, error)
	CreateUserService(userCreate entity.CreateUserRequest) (entity.User, error)
	UpdateUserService(id int, userUp entity.EditUserRequest) (entity.User, error)
	DeleteUserService(id int) (entity.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetUsersService() ([]entity.User, error) {
	users, err := s.repository.GetAllUser()
	return users, err
}

func (s *userService) CreateUserService(userCreate entity.CreateUserRequest) (entity.User, error) {
	user := entity.User{}
	user.Name = userCreate.Name
	user.Email = userCreate.Email
	user.Password = userCreate.Password
	user.Address = userCreate.Address

	createUser, err := s.repository.CreateUser(user)
	return createUser, err
}

func (s *userService) GetUserByIdService(id int) (entity.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func (s *userService) GetUserByNameService(name string) (entity.User, error) {
	user, err := s.repository.GetIdByName(name)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) UpdateUserService(id int, userUp entity.EditUserRequest) (entity.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return user, err
	}

	user.Name = userUp.Name
	user.Email = userUp.Email
	user.Password = userUp.Password
	user.Address = userUp.Address

	updateUser, err := s.repository.UpdateUser(user)
	return updateUser, err
}

func (s *userService) DeleteUserService(id int) (entity.User, error) {
	userID, err := s.GetUserByIdService(id)
	if err != nil {
		return userID, err
	}

	deleteUser, err := s.repository.DeleteUser(userID)
	return deleteUser, err
}
