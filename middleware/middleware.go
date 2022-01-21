package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTService interface {
	GenerateToken(username string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("iya")

func AuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userName"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}

func ExtractTokenUserName(jwtSecret string, e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userName := claims["userName"].(string)
		if userName == "" {
			return userName, fmt.Errorf("empty username")
		}
		return userName, nil
	}
	return "", fmt.Errorf("invalid user")
}
