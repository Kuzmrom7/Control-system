package main

import (
	"net/http"
	"time"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"

	"./oauth"
	"./handlersfunc"
)

var handler = http.HandlerFunc( func (w http.ResponseWriter, r *http.Request){

		w.Write([]byte("Hello, please change correct url!\n "))
})

func main(){
	r := mux.NewRouter()

	//ROUTES
	r.Handle("/", handler)
	r.Handle("/c/create", jwtMiddleware.Handler(handlersfunc.CreateContainer)).Methods("POST")
	r.Handle("/c/run", jwtMiddleware.Handler(handlersfunc.RunContainer)).Methods("GET")
	r.Handle("/c/stop", jwtMiddleware.Handler(handlersfunc.StopContainer)).Methods("GET")
	r.Handle("/c/delete", jwtMiddleware.Handler(handlersfunc.DeleteContainer)).Methods("DELETE")
	r.Handle("/c/info", jwtMiddleware.Handler(handlersfunc.InfoContainer)).Methods("GET")
	r.Handle("/c/list", jwtMiddleware.Handler(handlersfunc.ListContainer)).Methods("GET")

	r.HandleFunc("/gettoken", oauth.GetTokenHandler).Methods("GET")

	//Up and RUN server

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout,r),
		Addr:         "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}


//
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter:func(token *jwt.Token)(interface{}, error){
		return oauth.MySigninToken, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
