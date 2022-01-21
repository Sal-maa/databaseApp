package router

import (
	"database/sql"
	"designpattern/handler"
	"designpattern/middleware"
	"designpattern/repository"
	"designpattern/service"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func LogElapsedTime(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()
		err := handler(c)
		elapsed := time.Since(startTime)
		fmt.Println(elapsed)
		return err
	}
}

func UserRouter(e *echo.Echo, db *sql.DB, jwtSecret string) {
	e.Pre(echoMiddleware.RemoveTrailingSlash(), echoMiddleware.Logger())
	authService := middleware.AuthService()
	// Route User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(authService, userService)

	e.POST("/login", userHandler.AuthController)
	e.GET("/users", middleware.AuthMiddleware(authService, userService, userHandler.GetUsersController))
	e.POST("/users", userHandler.CreateUserController)
	e.GET("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.GetUserController))
	e.PUT("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.UpdateUserController))
	e.DELETE("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.DeleteUserController))

	// Route Book
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	e.GET("/books", bookHandler.GetBooksController)
	e.POST("/books", middleware.AuthMiddleware(authService, userService, bookHandler.CreateBookController))
	e.GET("/books/:id", bookHandler.GetBookController)
	e.PUT("/books/:id", middleware.AuthMiddleware(authService, userService, bookHandler.UpdateBookController))
	e.DELETE("/books/:id", middleware.AuthMiddleware(authService, userService, bookHandler.DeleteBookController))

	// Route Product
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	e.GET("/products", productHandler.GetProductsController)
	e.POST("/products", middleware.AuthMiddleware(authService, userService, productHandler.CreateProductController))
	e.GET("/products/:id", productHandler.GetProductController)
	e.PUT("/products/:id", middleware.AuthMiddleware(authService, userService, productHandler.UpdateProductController))
	e.DELETE("/products/:id", middleware.AuthMiddleware(authService, userService, productHandler.DeleteProductController))
}
