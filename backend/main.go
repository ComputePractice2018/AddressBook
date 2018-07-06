package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/AddressBook/backend/server"
)

func main() {
	//name := flag.String("name", "Иван", "имя для преветствия")
	//flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/api/addressbook/contacts", server.GetContacts).Methods("GET")
	router.HandleFunc("/api/addressbook/contacts", server.AddContact).Methods("POST")
	router.HandleFunc("/api/addressbook/contacts/{id}", server.EditContact).Methods("PUT")
	router.HandleFunc("/api/addressbook/contacts/{id}", server.DeleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
