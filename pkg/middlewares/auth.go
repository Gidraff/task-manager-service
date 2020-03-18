package middlewares

//
// import (
// 	"context"
// 	"net/http"
// 	"os"
// 	"strings"
//
// 	"github.com/Gidraff/task-manager-service/cmd/httpd/domain"
// 	"github.com/Gidraff/task-manager-service/cmd/httpd/httputil"
// 	jwt "github.com/dgrijalva/jwt-go"
// )
//
// // JwtAuthentication middleware
// var JwtAuthentication = func(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// List of endpoints that don't require auth
// 		noAuth := []string{"/api/user/new", "/api/user/login"}
// 		requestPath := r.URL.Path // Current request path
//
// 		for _, value := range noAuth {
// 			if value == requestPath {
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 		}
//
// 		response := make(map[string]interface{})
// 		tokenHeader := r.Header.Get("Authorization") // Grab token fron the header
//
// 		if tokenHeader == "" {
// 			response = httputil.Message(false, "Missing authorization token")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Add("Content-Type", "application/json")
// 			httputil.Respond(w, response)
// 			return
// 		}
//
// 		splitted := strings.Split(tokenHeader, " ")
// 		if len(splitted) != 2 {
// 			response = httputil.Message(false, "Invalid auth token")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Set("Content-Type", "application/json")
// 			httputil.Respond(w, response)
// 			return
// 		}
//
// 		extractedToken := splitted[1] // Extract token substring
// 		tk := &domain.Token{}
//
// 		token, err := jwt.ParseWithClaims(extractedToken, tk, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
// 		})
//
// 		if err != nil {
// 			response = httputil.Message(false, "Malformed authentication token")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Set("Content-Type", "application/json")
// 			httputil.Respond(w, response)
// 			return
// 		}
//
// 		if !token.Valid {
// 			response = httputil.Message(false, "Token is not valid")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Set("Content-Type", "application/json")
// 			httputil.Respond(w, response)
// 			return
// 		}
//
// 		// fmt.Sprintf("User %", tk.Username)
// 		ctx := context.WithValue(r.Context(), "user", tk.UserId)
// 		r = r.WithContext(ctx)
// 		next.ServeHTTP(w, r) //proceed in the middleware chain
// 	})
// }
