package models

type ContactInformation struct {

}

type Client struct {
	FirstName		string	    `json:"firstName"`
	LastName		string	    `json:"lastName"`
	ContactInfo interface{} `json:"contactInfo"`
}