package contact

import (
	"errors"
)

// Repo provides an abstraction layer over the DB
type Repo struct {
	db  *gorm.DB
}

// NewRepo creates a new Repo
func NewRepo(db *gorm.DB) (*Repo, error) {
	if db == nil {
		return nil, errors.New("db cannot be empty")
	}

	return &Repo{
		db:      db,
	}, nil
}

// GetAllContacts
func (r Repo) GetAllContacts() ([]*Contact, error) {
	var contacts []*Contact
	if err := r.Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

// CreateContact
func (r Repo) CreateContact(contact Contact) (*Contact, error) {
	var c Contact
	if err := r.Create(&contact).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

// GetContactById
func (r Repo) GetContactByID(ID string) (*Contact, error) {
	var contact Contact
	if err := r.Where("id = ?", id).First(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

// UpdateContact
func (r Repo) UpdateContact(id string) (*Contact, error) {
	var contact Contact
	if err := r.Where("id = ?", id).Save(contact); err != nil {
		return nil, err
	}

	return &contact, nil
}

// DeleteContact
func (r Repo) DeleteContact(id string) (*Contact, error) {
	var contact Contact
	if err := r.Where("id = ?", id).Delete(contact); err != nil {
		return nil, err
	}

	return &contact, nil
}
