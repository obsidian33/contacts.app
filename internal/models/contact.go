package models

import "strings"

var contacts = []Contact{
	{ID: 2, First: "Carson", Last: "Gross", Phone: "123-456-7890", Email: "carson@example.comz"},
	{ID: 3, First: "", Last: "", Phone: "", Email: "joe@example2.com"},
	{ID: 5, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe@example.com"},
	{ID: 6, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe1@example.com"},
	{ID: 7, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe2@example.com"},
	{ID: 8, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe3@example.com"},
	{ID: 9, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe4@example.com"},
	{ID: 10, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe5@example.com"},
	{ID: 11, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe6@example.com"},
	{ID: 12, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe7@example.com"},
	{ID: 13, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe8@example.com"},
	{ID: 14, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe9@example.com"},
	{ID: 15, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe10@example.com"},
	{ID: 16, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe11@example.com"},
	{ID: 17, First: "Joe", Last: "Blow", Phone: "123-456-7890", Email: "joe12@example.com"},
	{ID: 18, First: "", Last: "", Phone: "", Email: "restexample1@example.com"},
	{ID: 19, First: "", Last: "", Phone: "", Email: "restexample2@example.com"},
}

type ContactModelInterface interface {
	All() ([]Contact, error)
	Search(query string) ([]Contact, error)
	Save(Contact) error
}

type Contact struct {
	ID     int
	First  string
	Last   string
	Phone  string
	Email  string
	Errors map[string]string
}

type ContactModel struct {
}

func (c *ContactModel) All() ([]Contact, error) {
	return contacts, nil
}

func (c *ContactModel) Search(query string) ([]Contact, error) {
	query = strings.ToLower(query)
	var results []Contact
	for _, contact := range contacts {
		match_first := strings.Contains(strings.ToLower(contact.First), query)
		match_last := strings.Contains(strings.ToLower(contact.Last), query)
		match_email := strings.Contains(strings.ToLower(contact.Email), query)
		match_phone := strings.Contains(strings.ToLower(contact.Phone), query)
		if match_first || match_last || match_email || match_phone {
			results = append(results, contact)
		}
	}

	return results, nil
}

func (c *ContactModel) Save(contact Contact) error {
	return nil
}
