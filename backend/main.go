package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/AddressBook/backend/data"
	"github.com/ComputePractice2018/AddressBook/backend/server"
)

func main() {
	//name := flag.String("name", "Иван", "имя для преветствия")
	//flag.Parse()
	contactList := data.NewContactList()
	router := server.NewRouter(contactList)

	log.Fatal(http.ListenAndServe(":8080", router))
}
