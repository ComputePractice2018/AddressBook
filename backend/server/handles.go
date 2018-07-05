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
		w.Header().Add("Content-Type", "plain/text; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON marshaling error: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		log.Printf("Hadler write error: %v", err)
	}
}
