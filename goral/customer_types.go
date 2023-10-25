package goral

type Customer struct {
	Name        string   `gorm:"column:Name;type:string" json:"name" swaggertype:"string"`               // Name of the customer
	PhoneNumber string   `gorm:"column:PhoneNumber;type:string" json:"phoneNumber" swaggertype:"string"` // Contact phone number
	Id          int      `gorm:"column:Id;type:int" json:"id" swaggertype:"string"`                      // Identifier for the object
	AdvisorId   *Advisor `gorm:"column:AdvisorId;type:int" json:"advisor_id" swaggertype:"string"`       // Advisor
}
