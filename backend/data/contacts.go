package data

import (
	"fmt"
)

// Contact структура для хранения записи адресной книги
type Contact struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Github  string `json:"github"`
}

// contacts хранимый список контактов
var contacts []Contact

// GetContacts возращает список контактов
func GetContacts() []Contact {
	return contacts
}

// AddContact добавляет контакт contact в конец списка и возращает id
func AddContact(contact Contact) int {
	id := len(contacts)
	contacts = append(contacts, contact)
	return id
}

// EditContact изменяет контакт c id на contact
func EditContact(contact Contact, id int) error {
	if id < 0 || id >= len(contacts) {
		return fmt.Errorf("incorrect ID")
	}
	contacts[id] = contact
	return nil
}

// RemoveContact удаляет контакт по id
func RemoveContact(id int) error {
	if id < 0 || id >= len(contacts) {
		return fmt.Errorf("incorrect ID")
	}
	contacts = append(contacts[:id], contacts[id+1:]...)
	return nil
}
