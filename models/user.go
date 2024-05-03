package models

import "gorm.io/gorm"

const (
	AdminRole = "admin"
	UserRole  = "user"
)

type User struct {
	gorm.Model

	Email     string `json:"email"; gorm:"column:email"`
	Password  string `json:"password"; gorm:"column:password"`
	FirstName string `json:"first_name"; gorm:"column:FirstName"`
	LastName  string `json:"last_name"; gorm:"column:LastName"`
	Role      string `json:"role"; gorm:"column:role"`
}

func (u *User) BeforeCreate(db *gorm.DB) {
	u.Role = UserRole
}
