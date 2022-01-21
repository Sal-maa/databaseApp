package service

import (
	"designpattern/entity"
	"fmt"
)

func (s *userService) LoginUserService(login entity.Login) (entity.User, error) {
	name := login.Name
	password := login.Password

	user, err := s.repository.Login(name)
	if user.Password != password {
		return user, fmt.Errorf("Wrong combination")
	}
	return user, err
}
