package server

import (
	"github.com/ComputePractice2018/AddressBook/backend/data"

	"github.com/gorilla/mux"
)

func NewRouter(contactList data.Editable) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/addressbook/contacts", GetContacts(contactList)).Methods("GET")
	router.HandleFunc("/api/addressbook/contacts", AddContact(contactList)).Methods("POST")
	router.HandleFunc("/api/addressbook/contacts/{id}", EditContact(contactList)).Methods("PUT")
	router.HandleFunc("/api/addressbook/contacts/{id}", DeleteContact(contactList)).Methods("DELETE")
	return router
}
