package tokens

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/steelthedev/irs/models"
)

func GenerateToken(user models.User) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	tokenExpiryKey := os.Getenv("TOKEN_EXPIRE")
	tokenExpiry, err := strconv.Atoi(tokenExpiryKey)
	if err != nil {
		slog.Info(err.Error())
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenExpiry)).Unix()
	claims["role"] = user.Role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func ExtractToken(ctx *gin.Context) string {

	// Ensure access token pass into query aren't read

	token := ctx.Query("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func TokenValid(ctx *gin.Context) error {
	secretKey := os.Getenv("SECRET_KEY")
	tokenString := ExtractToken(ctx)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractIdFromToken(ctx *gin.Context) (uint, error) {
	secretKey := os.Getenv("SECRET_KEY")
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userId"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
