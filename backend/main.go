package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/AddressBook/backend/server"
)

func main() {
	//name := flag.String("name", "Иван", "имя для преветствия")
	//flag.Parse()

	http.HandleFunc("/api/addressbook/contacts", server.GetContacts)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
