package contact

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Getter
type Getter interface {
	GetAllContacts() ([]*Contact, error)
	GetContactByID(contactID string) (*Contact, error)
}

// Updater
type Updater interface {
	CreateContact(contact Contact) (*Contact, error)
	UpdateContact(contactID string) (*Contact, error)
	DeleteContact(contactID string) (*Contact, error)
}

type Response struct {
	IsSuccess bool
	Message   string
	Data      interface{}
}

//GetAllContacts
func GetAllContacts(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var apiResponse Response
		contacts, err := repo.GetAllContacts()
		if err != nil {
			JSON(w, http.StatusNotFound, apiResponse)
			return
		} else {
			apiResponse.IsSuccess = true
			apiResponse.Message = "Success"
			apiResponse.Data = contacts
			JSON(w, http.StatusOK, apiResponse)
			return
		}
	}
}

func GetContactByID(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var apiResponse Response
		id := r.Form.Get("id")
		contact, err := repo.GetContactByID(id)
		if err != nil {
			// apiResponse.IsSuccess = false
			apiResponse.Message = err.Error()
			JSON(w, http.StatusNotFound, apiResponse)
			return
		} else {
			apiResponse.IsSuccess = true
			apiResponse.Message = "Success"
			apiResponse.Data = contact
			JSON(w, http.StatusOK, apiResponse)
			return
		}
	}
}

func AddContact(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//c.BindJSON(&contact)
		var apiResponse Response
		contact, err := repo.CreateContact(Contact{})
		if err != nil {
			fmt.Printf("Error creating user: %s", err.Error())
			apiResponse.Message = err.Error()
			JSON(w, http.StatusInternalServerError, apiResponse)
			return
		} else {
			apiResponse.IsSuccess = true
			apiResponse.Message = "Success"
			apiResponse.Data = contact
			if err = JSON(w, http.StatusOK, apiResponse); err != nil {
				fmt.Println(fmt.Errorf("could not return success response %+v", err.Error()))
			}
			JSON(w, http.StatusOK, apiResponse)
			return
		}
	}
}

func UpdateContact(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Form.Get("id")
		var apiResponse Response
		contact, err := repo.GetContactByID(id)
		if err != nil {
			apiResponse.Message = err.Error()
			JSON(w, http.StatusNotFound, apiResponse)
			return
		} else {
			// update c
			fmt.Println(contact)
			contact, err := repo.UpdateContact(id)
			if err != nil {
				apiResponse.Message = err.Error()
				JSON(w, http.StatusInternalServerError, apiResponse)
				return
			}
			apiResponse.IsSuccess = true
			apiResponse.Message = "contact with id: " + id + " updated successfully"
			apiResponse.Data = contact
			JSON(w, http.StatusOK, apiResponse)
			return
		}
	}
}

func DeleteContact(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PostForm.Get("id")
		var apiResponse= new(Response)
		contact, err := repo.GetContactByID(id)
		if err != nil {
			apiResponse.Message = err.Error()
			JSON(w, http.StatusNotFound, apiResponse)
			return
		}

		contact, err = repo.DeleteContact(id)
		if err != nil {
			apiResponse.Message = err.Error()
			JSON(w, http.StatusInternalServerError, apiResponse)
			return
		}
		apiResponse.IsSuccess = true
		apiResponse.Message = "contact with id: " + id + " deleted successfully"
		apiResponse.Data = contact
		JSON(w, http.StatusOK, apiResponse)
		return
	}
}

// JSON render a generic interface as response of type json
func JSON(w http.ResponseWriter, code int, v interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if v == nil || code == http.StatusNoContent {
		return
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	if err = enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}