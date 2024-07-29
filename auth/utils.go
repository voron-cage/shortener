package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func getHashPassword(password, salt string) string {
	hash := password + salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(hash), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func compareHashPassword(password, salt, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt)) == nil
}

func obtainJWTToken(email, secretKey string) (string, error) {
	payload := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
