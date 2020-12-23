package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns a hashed password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

// ComparePassword compares password and return boolean values
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken returns token
func GenerateToken(userID uint64) (string, error) {
	// Creating access token
	atclaims := jwt.MapClaims{}
	atclaims["authorized"] = true
	atclaims["user_id"] = userID
	atclaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atclaims)
	token, err := at.SignedString([]byte("jdnfksdmfksd"))
	if err != nil {
		return "", err
	}

	return token, nil
}

// VerifyToken verify is a token match standard
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// verify token conforms to "SigningMethodHS256" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("jdnfksdmfksd"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken retrieves token from the Request
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// ExtractTokenMetadata returns auth id
func ExtractTokenMetadata(r *http.Request) (uint64, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}
	return 0, err
}
