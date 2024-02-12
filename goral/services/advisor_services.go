package goral

import (
	api_structure "generate/goral/structures"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAdvisor interface {
	GetAdvisor(filter api_structure.Advisor) ([]api_structure.Advisor, error)
	PaginateAdvisor(filter api_structure.Advisor) ([]api_structure.Advisor, error)
	InfiniteAdvisor(filter api_structure.Advisor) ([]api_structure.Advisor, error)
	CreateAdvisor(data api_structure.Advisor) (api_structure.Advisor, error)
	UpdateAdvisor(id int, data api_structure.Advisor) error
	DeleteAdvisor(id int) error
}
type AdvisorService struct{ DB *gorm.DB }

func (c *AdvisorService) GetAdvisor(filter api_structure.Advisor) ([]api_structure.Advisor, error) {
	result := []api_structure.Advisor{}
	var err error
	if err = c.DB.Preload(clause.Associations).Model(&api_structure.Advisor{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}

func (c *AdvisorService) PaginateAdvisor(filter api_structure.Advisor) ([]api_structure.Advisor, error) {
	return []api_structure.Advisor{}, nil
}

func (c *AdvisorService) InfiniteAdvisor(filter api_structure.Advisor) ([]api_structure.Advisor, error) {
	return []api_structure.Advisor{}, nil
}

func (c *AdvisorService) CreateAdvisor(data api_structure.Advisor) (api_structure.Advisor, error) {
	var err error
	if err = c.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}

func (c *AdvisorService) CreateBulkAdvisor(data []api_structure.Advisor) ([]api_structure.Advisor, error) {
	var err error
	if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil {
		return data, err
	}
	return data, err
}

func (c *AdvisorService) UpdateAdvisor(id int, data api_structure.Advisor) error {

	var err error
	if err = c.DB.Model(api_structure.Advisor{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}

func (c *AdvisorService) DeleteAdvisor(id int) error {
	var err error
	if err = c.DB.Where("id = ?", id).Delete(&api_structure.Advisor{}).Error; err != nil {
		return err
	}
	return err
}

func (c *AdvisorService) DeleteBulkAdvisor(ids []int) error {
	var err error
	for _, id := range ids {
		if err = c.DB.Where("id = ?", id).Delete(&api_structure.Advisor{}).Error; err != nil {
			return err
		}
	}
	return err
}
