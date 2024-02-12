package goral

import (
        "gorm.io/gorm" 
		"gorm.io/gorm/clause"
		api_structure "generate/goral/structures"
    )

 type ICustomer interface {
GetCustomer(filter api_structure.Customer) ([]api_structure.Customer, error) 
PaginateCustomer(filter api_structure.Customer) ([]api_structure.Customer, error) 
InfiniteCustomer(filter api_structure.Customer) ([]api_structure.Customer, error) 
CreateCustomer(data api_structure.Customer) (api_structure.Customer, error) 
UpdateCustomer(id int, data api_structure.Customer) error
DeleteCustomer(id int) error 
}
type CustomerService struct{ DB *gorm.DB }
func (c *CustomerService) GetCustomer(filter api_structure.Customer) ([]api_structure.Customer, error) { 
result := []api_structure.Customer{} 
var err error 
if err = c.DB.Preload(clause.Associations).Model(&api_structure.Customer{}).Where(filter).Find(&result).Error; err != nil { 
return result, err 
}
return result, err} 

func (c *CustomerService) PaginateCustomer(filter api_structure.Customer) ([]api_structure.Customer, error) {return []api_structure.Customer{},nil 
}

func (c *CustomerService) InfiniteCustomer(filter api_structure.Customer) ([]api_structure.Customer, error) {return []api_structure.Customer{},nil 
}

func (c *CustomerService) CreateCustomer(data api_structure.Customer) (api_structure.Customer, error) { 
var err error 
if err =  c.DB.Create(&data).Error; err != nil { 
return data, err 
}
return data, err} 

func (c *CustomerService) CreateBulkCustomer(data []api_structure.Customer) ([]api_structure.Customer, error) { 
var err error 
if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil { 
return data, err 
}
return data, err} 

func (c *CustomerService) UpdateCustomer(id int, data api_structure.Customer) error {

var err error 
if err = c.DB.Model(api_structure.Customer{}).Where("id = ?", id).Updates(&data).Error; err != nil {
return err
}
return err} 

func (c *CustomerService) DeleteCustomer(id int) error { 
var err error 
if err = c.DB.Where("id = ?", id).Delete(&api_structure.Customer{}).Error; err != nil { 
return err 
}
return err} 

func (c *CustomerService) DeleteBulkCustomer(ids []int) error {
var err error 
for _, id := range ids {
if err = c.DB.Where("id = ?", id).Delete(&api_structure.Customer{}).Error; err != nil {
return err
}
}
return err} 

