package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"gitlab.devkeeper.com/authreg/server/internal/app/respond"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("authorization")

		fmt.Printf(auth)

		if auth != "" {
			tokenString := auth[7:len(auth)]
			if tokenString != "" {
				saltKey := os.Getenv("AUTH_TOKEN_SALT_KEY")
				token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(saltKey), nil
				})
				if err == nil && token.Valid {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		response, _ := json.Marshal(map[string]string{"message": "not authorized"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write(response)
		return

	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user User
	var tokenString string
	var result string
	var message string

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if user.Login == os.Getenv("AUTH_LOGIN") && user.Password == os.Getenv("AUTH_PASSWORD") {

		result = "success"

		saltKey := os.Getenv("AUTH_TOKEN_SALT_KEY")

		/* type MyCustomClaims struct {
			Foo string `json:"foo"`
			jwt.StandardClaims
		} */

		claims := &jwt.StandardClaims{
			Issuer: "admin",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ = token.SignedString([]byte(saltKey))

	} else {
		result = "failed"
		message = "error validate"
	}

	respond.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"result":  result,
		"message": message,
		"token":   tokenString,
	})

}
