package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ComputePractice2018/AddressBook/backend/data"
	"github.com/gorilla/mux"
)

//GetContacts обрабатывает запросы на получение списка контактов
func GetContacts(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(cl.GetContacts())
		if err != nil {
			message := fmt.Sprintf("Unable to encode data: %v", err)
			http.Error(w, message, http.StatusInternalServerError)
			log.Println(message)
			return
		}
	}
}

//AddContact обрабатывает POST запрос
func AddContact(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var contact data.Contact
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			message := fmt.Sprintf("Unable to decode POST data: %v", err)
			http.Error(w, message, http.StatusUnsupportedMediaType)
			log.Println(message)
			return
		}
		id := cl.AddContact(contact)
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditContact обрабатывает PUT запрос
func EditContact(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var contact data.Contact
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			message := fmt.Sprintf("Unable to decode PUT data: %v", err)
			http.Error(w, message, http.StatusUnsupportedMediaType)
			log.Println(message)
			return
		}
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}

		err = cl.EditContact(contact, id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)

	}
}

// DeleteContact обрабатывает DELETE запрос
func DeleteContact(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}

		err = cl.RemoveContact(id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusNoContent)
	}
}
