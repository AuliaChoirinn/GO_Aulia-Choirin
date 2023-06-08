package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"Praktikum/configs"
	"Praktikum/models"
)

func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := configs.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	bookId := c.Param("id")
	var book models.Book

	if err := configs.DB.Where("ID = ?", bookId).Take(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := configs.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	bookId := c.Param("id")
	var book models.Book

	if err := configs.DB.Where("ID = ?", bookId).Take(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "book not found",
		})
	}

	if err := configs.DB.Delete(&models.Book{}, bookId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

func UpdateBookController(c echo.Context) error {
	var book models.Book
	bookId := c.Param("id")

	if err := configs.DB.Where("ID = ?", bookId).Take(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "book not found",
		})
	}

	var bookBody models.Book
	err := c.Bind(&bookBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	err = configs.DB.Where("ID = ?", bookId).Updates(models.Book{Title: bookBody.Title, Author: bookBody.Author, Description: bookBody.Description}).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err,
		})
	}

	book.Title = bookBody.Title
	book.Author = bookBody.Author
	book.Description = bookBody.Description

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}