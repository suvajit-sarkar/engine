package auth

import (
	"context"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/suvajit-sarkar/engine/rstorage"
)

var ctx = context.Background()

type Token struct {
	ID   string
	User User
}

type User struct {
	ID   string
	Name string
}

// Authenticate the user login function
func Authenticate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("userName")
	password := r.FormValue("userPassword")

	rauth := rstorage.NewAuth()
	userCred, error := rauth.FindUser(ctx, name)

	if len(name) == 0 || len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide name and password to obtain the token"))
		return
	}
	if error == nil && password == userCred.UserPassword {
		token, err := getToken(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			w.Header().Set("Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Token: " + token))

		}
	} else {
		print("Error resp")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Name and password do not match"))
		return
	}
}

// Middleware authenticates the users and passes it to the calling api
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Token:", "Bearer", 1)
		//TODO Check if it is an actual request for websocket
		if len(tokenString) == 0 {
			//TODO check null pointer exceptions
			vars := mux.Vars(r)
			tokenString = strings.Replace(vars["token"], "Token:", "Bearer", 1)
		}
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("This is not the page your looking for.."))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("name", name)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}
