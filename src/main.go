package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"./handler"
)

func main() {
	router := httprouter.New()
	registerDefaultHandler(router)
	handler.RegisterMessageHandler(router)
	http.ListenAndServe(":8080", router)
}

func registerDefaultHandler(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("welcome"))
	})
	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("pong"))
	})
}
