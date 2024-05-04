package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailIsValid(t *testing.T) {
	emails := []string{"akinwumikaliyanu@gmail.com", "AkinWumi2450-@gmail.com"}
	for _, email := range emails {
		emailValid := EmailIsValid(email)
		assert.True(t, emailValid, "Expecting all emails to pass")
	}
}

func TestEmailIsNotValid(t *testing.T) {
	emails := []string{"akinwumikal-iyanu2&@gmail.com", "akin8&%@gmail.com"}
	for _, email := range emails {
		emailValid := EmailIsValid(email)
		assert.False(t, emailValid, "Expecting all emails to fail")
	}
}

func TestHashPasswordSuccess(t *testing.T) {
	password := "passwordtobehashed"

	hashedPwd, err := HashPassword(password)
	assert.IsType(t, []byte{}, hashedPwd, "Expecting hashed password to be byte")
	assert.Nil(t, err, "Expecting errors to be nil")
}
