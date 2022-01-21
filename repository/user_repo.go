package repository

import (
	"database/sql"
	"designpattern/entity"
	"fmt"
)

type UserRepository interface {
	Login(name string) (entity.User, error)
	GetAllUser() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	GetUser(idParam int) (entity.User, error)
	GetIdByName(name string) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUser() ([]entity.User, error) {
	var users []entity.User

	result, err := r.db.Query("SELECT id,name,email,password,address FROM users")
	if err != nil {
		fmt.Println("failed in query", err)
	}
	defer result.Close()

	for result.Next() {
		var user entity.User
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Address)
		if err != nil {
			fmt.Println("failed to scan", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	_, err := r.db.Exec("INSERT INTO users(name, email, password, address,phone) VALUES(?,?,?,?,?)", user.Name, user.Email, user.Password, user.Address, user.Phone)
	return user, err
}

func (r *userRepository) GetUser(idParam int) (entity.User, error) {
	var user entity.User
	result, err := r.db.Query("SELECT id, name, email, password, address FROM users WHERE id=?", idParam)
	if err != nil {
		fmt.Println("failed in query", err)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}

	errScan := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Address)
	if errScan != nil {
		fmt.Println("failed to read data", err)
	}

	if idParam == user.Id {
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

func (r *userRepository) GetIdByName(name string) (entity.User, error) {
	var user entity.User
	result, err := r.db.Query("SELECT id,name FROM users WHERE name=?", name)
	if err != nil {
		fmt.Println("failed in query", err)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}

	errScan := result.Scan(&user.Id, &user.Name)
	if errScan != nil {
		fmt.Println("failed to read data", errScan)
	}

	if name == user.Name {
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

func (r *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	result, err := r.db.Exec("UPDATE users SET name=?, email=?, password=?, address=? WHERE id=?", user.Name, user.Email, user.Password, user.Address, user.Id)
	if err != nil {
		return user, fmt.Errorf("failed to update data", err)
	}
	NotAffected, _ := result.RowsAffected()
	if NotAffected == 0 {
		return user, fmt.Errorf("failed to find data id")
	}

	return user, nil
}

func (r *userRepository) DeleteUser(user entity.User) (entity.User, error) {
	_, err := r.db.Exec("DELETE FROM users WHERE id=?", user.Id)
	return user, err
}
