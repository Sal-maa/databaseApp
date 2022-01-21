package entity

type Login struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Token string `json:"token"`
}
