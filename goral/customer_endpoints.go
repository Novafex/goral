package goral

import "github.com/gofiber/fiber/v2"

type CustomerController struct{ Svc CustomerService }

func (controller *CustomerController) GetCustomer(c *fiber.Ctx) error         { return nil }
func (controller *CustomerController) PaginateCustomer(c *fiber.Ctx) error    { return nil }
func (controller *CustomerController) InfiniteCustomer(c *fiber.Ctx) error    { return nil }
func (controller *CustomerController) CreateCustomer(c *fiber.Ctx) error      { return nil }
func (controller *CustomerController) Create_bulkCustomer(c *fiber.Ctx) error { return nil }
func (controller *CustomerController) UpdateCustomer(c *fiber.Ctx) error      { return nil }
func (controller *CustomerController) DeleteCustomer(c *fiber.Ctx) error      { return nil }
func (controller *CustomerController) Delete_bulkCustomer(c *fiber.Ctx) error { return nil }
