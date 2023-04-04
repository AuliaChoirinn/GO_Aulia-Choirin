package lib

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

const (
	SecretKey       = "my_secret_key"
	TokenExpiration = time.Hour * 2
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpiration).Unix(),
		},
	})

	return token.SignedString([]byte(SecretKey))
}

func ExtractTokenUserID(c echo.Context) (uint, error) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(*Token)
		return claims.UserID, nil
	}
	return 0, fmt.Errorf("invalid token")
}
