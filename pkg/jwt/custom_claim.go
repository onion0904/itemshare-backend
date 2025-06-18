package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// ログイン認証用のclaim
type MyCustomClaims struct {
	Email  string `json:"email"`
	UserID string `json:"UserID"`
	jwt.RegisteredClaims
}

func NewCustomClaims(email string, userID string) *MyCustomClaims {
	claims := MyCustomClaims{
		Email:  email,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 有効期限
		},
	}
	return &claims
}

// group招待用claim
type InviteClaims struct {
	GroupID string `json:"group_id"`
	jwt.RegisteredClaims
}

func NewInviteClaims(groupID string) *InviteClaims {
	claims := InviteClaims{
		GroupID: groupID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	return &claims
}
