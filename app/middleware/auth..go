package middleware

// 参照元 https://zenn.dev/hsaki/books/golang-graphql/viewer/auth

import (
	"fmt"
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/onion0904/CarShareSystem/app/config"
	"github.com/onion0904/CarShareSystem/pkg/jwt"
)

type userIDKey struct{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, req)
			return
		}

		userID, err := validateToken(token)
		if err != nil {
			log.Println(err)
			http.Error(w, `{"reason": "invalid token"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), userIDKey{}, userID)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func GetUserID(ctx context.Context) (string, bool) {
	switch v := ctx.Value(userIDKey{}).(type) {
	case string:
		return v, true
	default:
		return "", false
	}
}

func validateToken(token string) (string, error) {
	config := config.GetConfig()
	secretkey := config.JWT.Secret
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	fmt.Println("Received Token:", token)
	claims, err := jwt.ParseJWT(token,[]byte(secretkey))
	if err != nil {
		return "", err
	}

	return claims.UserID, nil
}