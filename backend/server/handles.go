package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/AddressBook/backend/data"
)

//ContactsHandler обрабатывает все запросы к /api/addressbook/contacts
func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetContacts(w, r)
		return
	}
	if r.Method == "POST" {
		AddContact(w, r)
		return
	}

	message := fmt.Sprintf("Method %s is not allowed", r.Method)
	http.Error(w, message, http.StatusMethodNotAllowed)
	log.Println(message)
}

//GetContacts обрабатывает запросы на получение списка контактов
func GetContacts(w http.ResponseWriter, r *http.Request) {
	binaryData, err := json.Marshal(data.ContactList)
	if err != nil {
		message := fmt.Sprintf("JSON marshaling error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		message := fmt.Sprintf("Handler write error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
	}
}

//AddContact обрабатывает POST запрос
func AddContact(w http.ResponseWriter, r *http.Request) {
	var contact data.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	log.Printf("%+v", contact)
	w.WriteHeader(http.StatusCreated)
}
