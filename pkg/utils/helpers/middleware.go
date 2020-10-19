package helpers

import (
	"context"
	//"fmt"
	"log"
	"net/http"
	//"strings"
	//"github.com/dgrijalva/jwt-go"
)

// JwtMiddleware middleware
func JwtMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var msg string
	// List of endpoints that don't require auth
	noAuth := []string{"/api/v1/auth/sign-in", "/api/v1/auth/sign-up"}
	requestPath := r.URL.Path // Current request path

	for _, value := range noAuth {
		if value == requestPath {
			log.Printf("public path %s", value)
			next(w, r)
			return
		}
	}

	tokenHeader := r.Header.Get("Authorization")
	if len(tokenHeader) == 0 {
		msg = "Missing authorization token"
		Response(http.StatusUnauthorized, msg, w)
		return
	}

	token, err := VerifyToken(r)
	if err != nil {
		Response(http.StatusUnauthorized, "Invalid authentication token", w)
		return
	}

	if !token.Valid {
		msg = "Token is not valid"
		Response(http.StatusUnauthorized, msg, w)
		return
	}

	// fmt.Sprintf("User %", tk.Username)
	ctx := context.WithValue(r.Context(), "user", token.Claims)
	next(w, r.WithContext(ctx)) //proceed in the middleware chain
}
