package goral

type Customer struct {
	Id int `gorm:"column:Id;type:int" json:"id" swaggertype:"string"`	// Identifier for the object
	AdvisorId *Advisor `gorm:"column:AdvisorId;type:nullable" json:"advisor_id" swaggertype:"string"`	// Advisor
	Name string `gorm:"column:Name;type:string" json:"name" swaggertype:"string"`	// Name of the customer
	PhoneNumber string `gorm:"column:PhoneNumber;type:string" json:"phoneNumber" swaggertype:"string"`	// Contact phone number
}
