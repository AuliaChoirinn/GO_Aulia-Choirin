package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.JWT([]byte("secret")))

    // Routes for User model
    e.GET("/users", GetUsersController)
    e.GET("/users/:id", GetUserController)
    e.POST("/users", CreateUserController)
    e.DELETE("/users/:id", DeleteUserController)

    // Routes for Book model
    e.GET("/books", GetBooksController)
    e.GET("/books/:id", GetBookController)
    e.POST("/books", CreateBookController)
    e.DELETE("/books/:id", DeleteBookController)
    e.PUT("/books/:id", UpdateBookController)

    e.Logger.Fatal(e.Start(":8080"))
}
