package oauth

import (
	"strings"
	"time"
	"net/http"
	"encoding/base64"

	"github.com/dgrijalva/jwt-go"
)
//Global key
var MySigninToken = []byte("secret")

var GetTokenHandler = http.HandlerFunc(func( w http.ResponseWriter, r * http.Request){

	auth := strings.SplitN(r.Header["Autorization"][0], " ", 2)
	if len(auth) != 2 || auth[0] != "Basic"{
		http.Error(w, "Bad syntax", http.StatusBadRequest)
		return
	}
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if Validate(pair[0], pair[1]){
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)

		claims["admin"] = true
		claims["name"] = "haipe"
		claims["exp"] = time.Now().Add(time.Hour *24).Unix()

		tokenString,err := token.SignedString(MySigninToken)
		if err != nil{
			panic(err)
		}

		w.Write([]byte(tokenString))
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}


})


//
func Validate(username, password string) bool {
	if username == "roman" && password == "kuzmenko" {
		return true
	}
	return false
}

