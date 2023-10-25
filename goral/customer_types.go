package goral

type Customer struct {
	PhoneNumber string `json:"phoneNumber,omitempty"` // Contact phone number
	Id          int    `json:"id"`                    // Identifier for the object
	Name        string `json:"name"`                  // Name of the customer
}
