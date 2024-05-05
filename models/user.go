package models

import (
	"log/slog"

	"gorm.io/gorm"
)

const (
	AdminRole = "admin"
	UserRole  = "user"
)

type User struct {
	gorm.Model

	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password" gorm:"column:password"`
	FirstName string `json:"first_name" gorm:"column:FirstName"`
	LastName  string `json:"last_name" gorm:"column:LastName"`
	Role      string `json:"role" gorm:"column:role"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.Role = UserRole

	return nil
}

func CheckUserExistsWithEmail(email string, db *gorm.DB) bool {
	var user User
	if result := db.Where("Email=?", email).First(&user); result.Error != nil {
		return false
	}
	return true
}

func GetUserById(ID uint, db *gorm.DB) (*User, error) {
	var user User
	if result := db.Where("ID=?", ID).First(&user); result.Error != nil {
		slog.Info(result.Error.Error())
		return nil, result.Error
	}
	return &user, nil
}

func GetUserRole(ID uint, db *gorm.DB) (string, error) {
	user, err := GetUserById(ID, db)
	if err != nil {
		slog.Info("User could not be fetched ", "Error", err.Error())
		return "", err
	}

	return user.Role, nil
}
