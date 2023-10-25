package goral

import "gorm.io/gorm"

type ICustomer interface {
	getCustomer(filter Customer) ([]Customer, error)
	paginateCustomer(filter Customer) ([]Customer, error)
	infiniteCustomer(filter Customer) ([]Customer, error)
	createCustomer(data Customer) (Customer, error)
	create_bulkCustomer(data []Customer) ([]Customer, error)
	updateCustomer(id int, data Customer) error
	deleteCustomer(id int) error
	delete_bulkCustomer(id int) error
}
type CustomerService struct{ DB *gorm.DB }

func (c *CustomerService) getCustomer(filter Customer) ([]Customer, error) {
	return []Customer{}, nil
}

func (c *CustomerService) paginateCustomer(filter Customer) ([]Customer, error) {
	return []Customer{}, nil
}

func (c *CustomerService) infiniteCustomer(filter Customer) ([]Customer, error) {
	return []Customer{}, nil
}

func (c *CustomerService) createCustomer(data Customer) (Customer, error) {
	return Customer{}, nil
}

func (c *CustomerService) create_bulkCustomer(data []Customer) ([]Customer, error) {
	return []Customer{}, nil
}

func (c *CustomerService) updateCustomer(id int, data Customer) error {
	return nil
}

func (c *CustomerService) deleteCustomer(id int) error {
	return nil
}

func (c *CustomerService) delete_bulkCustomer(id int) error {
	return nil
}
