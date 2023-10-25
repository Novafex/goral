package goral

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ICustomer interface {
	GetCustomer(filter Customer) ([]Customer, error)
	PaginateCustomer(filter Customer) ([]Customer, error)
	InfiniteCustomer(filter Customer) ([]Customer, error)
	CreateCustomer(data Customer) (Customer, error)
	UpdateCustomer(id int, data Customer) error
	DeleteCustomer(id int) error
}
type CustomerService struct{ DB *gorm.DB }

func (c *CustomerService) GetCustomer(filter Customer) ([]Customer, error) {
	result := []Customer{}
	var err error
	if err = c.DB.Preload(clause.Associations).Model(&Customer{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *CustomerService) PaginateCustomer(filter Customer) ([]Customer, error) {
	return []Customer{}, nil
}

func (c *CustomerService) InfiniteCustomer(filter Customer) ([]Customer, error) {
	return []Customer{}, nil
}

func (c *CustomerService) CreateCustomer(data Customer) (Customer, error) {
	var err error
	if err = c.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}

func (c *CustomerService) CreateBulkCustomer(data []Customer) ([]Customer, error) {
	var err error
	if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil {
		return data, err
	}
	return data, err
}

func (c *CustomerService) UpdateCustomer(id int, data Customer) error {

	var err error
	if err = c.DB.Model(Customer{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}

func (c *CustomerService) DeleteCustomer(id int) error {
	var err error
	if err = c.DB.Where("id = ?", id).Delete(&Customer{}).Error; err != nil {
		return err
	}
	return err
}

func (c *CustomerService) DeleteBulkCustomer(ids []int) error {
	var err error
	for _, id := range ids {
		if err = c.DB.Where("id = ?", id).Delete(&Customer{}).Error; err != nil {
			return err
		}
	}
	return err
}
