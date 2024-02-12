package goral

type Advisor struct {
	Id          int    `gorm:"column:Id;type:int" json:"id" swaggertype:"string"`                      // Identifier for the object
	Name        string `gorm:"column:Name;type:string" json:"name" swaggertype:"string"`               // Name of the advisor
	PhoneNumber string `gorm:"column:PhoneNumber;type:string" json:"phoneNumber" swaggertype:"string"` // Contact phone number
}
