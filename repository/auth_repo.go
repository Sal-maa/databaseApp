package repository

import (
	"designpattern/entity"
	"fmt"
)

func (r *userRepository) Login(name string) (entity.User, error) {
	user := entity.User{}
	result, err := r.db.Query("SELECT name,password FROM users WHERE name=? ", name)
	if err != nil {
		return user, err
	}
	defer result.Close()

	if isExist := result.Next(); !isExist {
		return user, fmt.Errorf("user not exist")
	}

	errScan := result.Scan(&user.Name, &user.Password)

	if errScan != nil {
		fmt.Println(errScan)
		return user, fmt.Errorf("error scanning data")
	}

	if user.Name == name {
		// usernya benar-benar ada
		return user, nil
	}
	// tidak error, tapi usernya tidak ada
	return user, fmt.Errorf("user not found")
}
