package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Name     string `json:"name" form:"name"`
    Email    string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
    var users []User
    err := db.Find(&users).Error
    if err != nil {
        return nil, err
    }
    return
