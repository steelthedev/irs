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

type UserService struct {
	DB *gorm.DB
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.Role = UserRole

	return nil
}

func (us UserService) CheckUserExistsWithEmail(email string) bool {
	var user User
	if result := us.DB.Where("Email=?", email).First(&user); result.Error != nil {
		return false
	}
	return true
}

func (us UserService) GetUserById(ID uint) (*User, error) {
	var user User
	if result := us.DB.Where("ID=?", ID).First(&user); result.Error != nil {
		slog.Info(result.Error.Error())
		return nil, result.Error
	}
	return &user, nil
}

func (us UserService) GetUserRole(ID uint) (string, error) {
	user, err := us.GetUserById(ID)
	if err != nil {
		slog.Info("User could not be fetched ", "Error", err.Error())
		return "", err
	}

	return user.Role, nil
}
