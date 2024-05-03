package utils

import (
	"log/slog"
	"regexp"

	"github.com/steelthedev/irs/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func EmailIsValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)

}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckUserExistsWithEmail(email string, db *gorm.DB) bool {
	var user models.User
	if result := db.Where("Email=?", email).First(&user); result.Error != nil {
		return false
	}
	return true
}

func GetUserById(ID uint, db *gorm.DB) (*models.User, error) {
	var user models.User
	if result := db.Where("ID=?", ID).First(&user); result.Error != nil {
		slog.Info(result.Error.Error())
		return nil, result.Error
	}
	return &user, nil
}
