package jwt


import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken (claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}