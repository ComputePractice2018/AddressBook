package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/AddressBook/backend/data"
)

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

func EditContact(w http.ResponseWriter, r *http.Request) {

}

func DeleteContact(w http.ResponseWriter, r *http.Request) {

}
