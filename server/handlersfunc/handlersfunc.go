package handlersfunc

import "net/http"


var CreateContainer = http.HandlerFunc( func (w http.ResponseWriter, r *http.Request){

})

var RunContainer = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

})

var StopContainer = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {

})

var DeleteContainer = http.HandlerFunc( func (w http.ResponseWriter, r *http.Request){

})

var InfoContainer = http.HandlerFunc (func (w http.ResponseWriter, r *http.Request)  {

})

var ListContainer = http.HandlerFunc ( func (w http.ResponseWriter, r *http.Request) {

})