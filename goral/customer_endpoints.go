package goral

import (
	"github.com/gofiber/fiber/v2"
)

type CustomerController struct{ Svc CustomerService }

// ShowCustomer godoc
// @Summary Show Customer
// @Description Get GetCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/get
func (controller *CustomerController) GetCustomer(c *fiber.Ctx) error { return nil }

// ShowCustomer godoc
// @Summary Show Customer
// @Description Paginate PaginateCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/paginate
func (controller *CustomerController) PaginateCustomer(c *fiber.Ctx) error { return nil }

// ShowCustomer godoc
// @Summary Show Customer
// @Description Infinite InfiniteCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/infinite
func (controller *CustomerController) InfiniteCustomer(c *fiber.Ctx) error { return nil }

// ShowCustomer godoc
// @Summary Show Customer
// @Description Create CreateCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/create
func (controller *CustomerController) CreateCustomer(c *fiber.Ctx) error { return nil }

// ShowCustomer godoc
// @Summary Show Customer
// @Description CreateBulk CreateBulkCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/createbulk
func (controller *CustomerController) CreateBulkCustomer(c *fiber.Ctx) error { return nil }

// ShowCustomer godoc
// @Summary Show Customer
// @Description Update UpdateCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/update
func (controller *CustomerController) UpdateCustomer(c *fiber.Ctx) error { return nil }

// ShowCustomer godoc
// @Summary Show Customer
// @Description Delete DeleteCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/delete
func (controller *CustomerController) DeleteCustomer(c *fiber.Ctx) error { return nil }

// ShowCustomer godoc
// @Summary Show Customer
// @Description DeleteBulk DeleteBulkCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/deletebulk
func (controller *CustomerController) DeleteBulkCustomer(c *fiber.Ctx) error { return nil }
