package auth

// import (
// 	"encoding/json"
// 	"net/http"
// 	"os"

// 	jwt "github.com/dgrijalva/jwt-go"
// )

// var SecretKey = []byte(os.Getenv("Cookie_Key"))

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, _ := r.Cookie("Session")
// 		tokenStr := cookie.Value
// 		claims := &Claims{}

// 		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
// 			// check token signing method etc
// 			return SecretKey, nil
// 		})

// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				http.Error(w, "Forbidden", http.StatusForbidden)
// 				return
// 			}
// 			w.WriteHeader(http.StatusBadRequest)
// 			http.Error(w, "Bad Request", http.StatusForbidden)
// 			return
// 		}

// 		if !token.Valid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			http.Error(w, "Forbidden", http.StatusForbidden)
// 			return
// 		}

// 		foundUser, err := database.GetUser(claims.Username)
// 		if err != nil {
// 			http.Error(w, "User Not Found", http.StatusForbidden)
// 			return
// 		}
// 		jsonResponse, err := json.Marshal("Authorised User: " + foundUser.Username)
// 		if err != nil {
// 			return
// 		}
// 		//update response
// 		w.Write(jsonResponse)
// 		next.ServeHTTP(w, r)
// 	})
// }
