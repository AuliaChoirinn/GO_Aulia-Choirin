package middleware

import (
	"fmt"
	"Go_Aulia-Choirin_Aulia-Choirin-Aulia-Choirin/22_UnitTesting/praktikum/config"
	"Go_Aulia-Choirin_Aulia-Choirin-Auliani-Choirin/22_UnitTesting/praktikum/models"

	"github.com/labstack/echo"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var user models.User
	fmt.Println(email, password)
	if err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
