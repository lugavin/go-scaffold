package main

import (
	"os"
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"./handler"
)

var (
	conf Config
)

type Config struct {
	Port int `json:"port"`
}

func init() {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/conf.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf = Config{}
	if err := decoder.Decode(&conf); err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := httprouter.New()
	registerDefaultHandler(router)
	handler.RegisterMessageHandler(router)
	port := strconv.Itoa(conf.Port)
	http.ListenAndServe(":"+port, router)
}

func registerDefaultHandler(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("welcome"))
	})
	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("pong"))
	})
}
