package main

import (
	"errors"
)

type Contact struct {
	name  string
	phone string
}

type ContactList map[string]Contact

func (cl *ContactList) listAll() []string {
	var contacts []string
	for k := range *cl {
		contacts = append(contacts, k)
	}
	return contacts
}

func (cl *ContactList) contactInfo(name string) (Contact, error) {
	contact, exists := (*cl)[name]
	if !exists {
		return contact, errors.New("contact not found")
	}
	return contact, nil
}

func (cl *ContactList) createContact(contact Contact) error {
	if _, exists := (*cl)[contact.name]; exists {
		return errors.New("contact already exists")
	}
	(*cl)[contact.name] = contact
	return nil
}

func (cl *ContactList) updateContact(name, phone string) error {
	if _, exists := (*cl)[name]; !exists {
		return errors.New("contact not found")
	}
	(*cl)[name] = Contact{name, phone}
	return nil
}

func (cl *ContactList) deleteContact(name string) error {
	if _, exists := (*cl)[name]; !exists {
		return errors.New("contact not found")
	}
	delete(*cl, name)
	return nil
}
