package data

import "testing"

var testContacts = []Contact{
	{
		ID:      1,
		Name:    "Ivan",
		Surname: "Ivanov",
		Phone:   "+7999999999",
		Email:   "iivanov@mail.com",
		Github:  "iivanov",
	},
	{
		ID:      2,
		Name:    "Sergey",
		Surname: "Sergeev",
		Phone:   "+788888888",
		Email:   "ssergeev@mail.com",
		Github:  "ssergeev",
	},
}

func TestAddContact(t *testing.T) {
	cl := NewContactList()

	cl.AddContact(testContacts[0])

	if cl.GetContacts()[0] != testContacts[0] {
		t.Errorf("AddContact is not working")
	}
}

func TestEditContact(t *testing.T) {
	cl := NewContactList()
	cl.AddContact(testContacts[0])

	err := cl.EditContact(testContacts[1], 1)

	if cl.GetContacts()[0] != testContacts[1] {
		t.Errorf("EditContact is not working")
	}
	if err != nil {
		t.Errorf("Unexpected EditContact error: %+v", err)
	}

	err = cl.EditContact(testContacts[1], -1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}

}

func TestRemoveContact(t *testing.T) {
	cl := NewContactList()
	cl.AddContact(testContacts[0])
	cl.AddContact(testContacts[1])

	err := cl.RemoveContact(1)

	if cl.GetContacts()[0] != testContacts[1] {
		t.Errorf("RemoveContact is not working")
	}
	if err != nil {
		t.Errorf("Unexpected RemoveContact error")
	}

	err = cl.RemoveContact(-1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}

}
