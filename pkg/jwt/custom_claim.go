package jwt

import (
	"github.com/golang-jwt/jwt/v5"
    "time"
)

type MyCustomClaims struct {
	Email string `json:"email"`
    UserID string `json:"UserID"`
	jwt.RegisteredClaims
}

func NewCustomClaims(email string, userID string) *MyCustomClaims {
	claims := MyCustomClaims{
        Email: email,
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            Subject:   email,
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 有効期限
        },
    }
	return &claims
}