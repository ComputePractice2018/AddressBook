package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ComputePractice2018/AddressBook/backend/data"
)

var testContact = "{\"name\":\"Сергей\",\"surname\":\"Сергеев\",\"phone\":\"+7-999-999-99-99\",\"email\":\"mail@domain.ru\",\"github\":\"user\"}"

func TestCrudHandlers(t *testing.T) {
	cl := data.NewContactList()
	router := NewRouter(cl)

	req, err := http.NewRequest("GET", "/api/addressbook/contacts", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}
	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}

	testData := strings.NewReader(testContact)
	req, err = http.NewRequest("POST", "/api/addressbook/contacts", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}
	if resp.Header.Get("Location") != "/api/addressbook/contacts/1" {
		t.Error("Expected another location")
	}

	if len(cl.GetContacts()) != 1 {
		t.Error("Expected new value")
	}

	testData = strings.NewReader(testContact)
	req, err = http.NewRequest("PUT", "/api/addressbook/contacts/1", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}
	if resp.Header.Get("Location") != "/api/addressbook/contacts/1" {
		t.Error("Expected another location")
	}

	if len(cl.GetContacts()) != 1 {
		t.Error("Expected old value")
	}

	req, err = http.NewRequest("DELETE", "/api/addressbook/contacts/1", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten: %d)", resp.StatusCode)
	}
	if len(cl.GetContacts()) != 0 {
		t.Error("Expected old value")
	}
}
