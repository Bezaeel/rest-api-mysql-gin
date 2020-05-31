package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/bezaeel/rest-api-mysql-gin/pkg/contact"
)


// initRoutes initialize the routing configuration and return a prepared http.Handler
func initRoutes(
	repo *contact.Repo,
) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Route("/contacts", func(r chi.Router) {
		r.Get("/", contact.GetAllContacts(repo))
	})

	r.Route("/contacts", func(r chi.Router) {
		r.Post("/", contact.AddContact(repo))
	})

	r.Route("/contacts/id", func(r chi.Router) {
		r.Get("/", contact.GetContactByID(repo))
	})

	r.Route("/contacts/id", func(r chi.Router) {
		r.Put("/", contact.UpdateContact(repo))
	})

	r.Route("/contacts/id", func(r chi.Router) {
		r.Delete("/", contact.DeleteContact(repo))
	})

	return r, nil
}