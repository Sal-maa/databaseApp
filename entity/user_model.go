package entity

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
}

type Address struct {
	Id      int    `json:"id" form:"id"`
	City    string `json:"city" form:"city"`
	Zipcode int    `json:"zipcode" form:"zipcode"`
	Street  string `json:"street" form:"street"`
}

type CreateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
}

type EditUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
}

type UserResponse struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}

func FormatUserResponse(user User) UserResponse {
	return UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}
}
