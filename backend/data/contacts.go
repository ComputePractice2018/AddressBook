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

// ContactList структрура для списка записей адресной книги
type ContactList struct {
	contacts []Contact
}

// Editable интерфейс для работы со списком записей
type Editable interface {
	GetContacts() []Contact
	AddContact(contact Contact) int
	EditContact(contact Contact, id int) error
	RemoveContact(id int) error
}

// NewContactList конструктор списка контактов
func NewContactList() *ContactList {
	return &ContactList{}
}

// GetContacts возращает список контактов
func (cl *ContactList) GetContacts() []Contact {
	return cl.contacts
}

// AddContact добавляет контакт contact в конец списка и возращает id
func (cl *ContactList) AddContact(contact Contact) int {
	id := len(cl.contacts)
	cl.contacts = append(cl.contacts, contact)
	return id
}

// EditContact изменяет контакт c id на contact
func (cl *ContactList) EditContact(contact Contact, id int) error {
	if id < 0 || id >= len(cl.contacts) {
		return fmt.Errorf("incorrect ID")
	}
	cl.contacts[id] = contact
	return nil
}

// RemoveContact удаляет контакт по id
func (cl *ContactList) RemoveContact(id int) error {
	if id < 0 || id >= len(cl.contacts) {
		return fmt.Errorf("incorrect ID")
	}
	cl.contacts = append(cl.contacts[:id], cl.contacts[id+1:]...)
	return nil
}
