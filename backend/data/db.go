package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DBContactList struct {
	db *gorm.DB
}

func NewDBContactList(connection string) (*DBContactList, error) {
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Contact{})
	return &DBContactList{db: db}, db.Error
}

func (cl *DBContactList) Close() {
	cl.db.Close()
}

func (cl *DBContactList) GetContacts() []Contact {
	var contacts []Contact
	cl.db.Find(&contacts)
	return contacts
}

func (cl *DBContactList) AddContact(contact Contact) int {
	cl.db.Create(&contact)
	return contact.ID
}

func (cl *DBContactList) EditContact(contact Contact, id int) error {
	var contacts []Contact
	cl.db.Where("id = ?", id).Find(&contacts)
	if len(contacts) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	contact.ID = contacts[0].ID
	cl.db.Save(&contact)
	return cl.db.Error
}

func (cl *DBContactList) RemoveContact(id int) error {
	var contacts []Contact
	cl.db.Where("id = ?", id).Find(&contacts)
	if len(contacts) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}
	cl.db.Delete(&contacts[0])
	return cl.db.Error
}
