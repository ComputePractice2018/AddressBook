package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/AddressBook/backend/data"
	"github.com/ComputePractice2018/AddressBook/backend/server"
)

func main() {
	//name := flag.String("name", "Иван", "имя для преветствия")
	//flag.Parse()
	contactList := data.NewContactList()
	router := mux.NewRouter()
	router.HandleFunc("/api/addressbook/contacts", server.GetContacts(contactList)).Methods("GET")
	router.HandleFunc("/api/addressbook/contacts", server.AddContact(contactList)).Methods("POST")
	router.HandleFunc("/api/addressbook/contacts/{id}", server.EditContact(contactList)).Methods("PUT")
	router.HandleFunc("/api/addressbook/contacts/{id}", server.DeleteContact(contactList)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
