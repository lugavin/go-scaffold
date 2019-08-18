package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"../model"
)

func RegisterMessageHandler(router *httprouter.Router) {
	router.POST("/api/messages", createMessage)
	router.PUT("/api/messages/:id", updateMessage)
	router.GET("/api/messages/:id", getMessage)
	router.DELETE("/api/messages/:id", deleteMessage)
}

func createMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var msg model.Message
	if json.Unmarshal(body, &msg) != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if model.InsertMessage(&msg) != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result, err := json.Marshal(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write(result)
	}
}

func updateMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var msg model.Message
	if json.Unmarshal(body, &msg) != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if model.UpdateMessage(ps.ByName("id"), &msg) != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if result, err := json.Marshal(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func getMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	message, err := model.GetMessage(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result, err := json.Marshal(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write(result)
	}
}

func deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := model.DeleteMessage(ps.ByName("id")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
