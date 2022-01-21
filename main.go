package main

import (
	"designpattern/entity"
	"designpattern/router"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := entity.InitDB(connectionString)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	router.UserRouter(e, db, jwtSecret)

	e.Logger.Fatal(e.Start(":8080"))

}
