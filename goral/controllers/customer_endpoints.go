package goral

import (
	api_service "generate/goral/services"
	api_structure "generate/goral/structures"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type CustomerController struct{ Svc api_service.CustomerService }

// ShowCustomer godoc
// @Summary Show Customer
// @Description Get GetCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} api_structure.Customer
// @Router /test/customer/get
func (controller *CustomerController) GetCustomer(c *fiber.Ctx) error {
	var customer []api_structure.Customer

	db := controller.Svc.DB.Table("Customer")
	query := c.Query("query")
	advisor_id := c.Query("advisor_id")
	db = db.Where("advisor_id = ?", advisor_id)

	finalQuery := "Name ILIKE" + "%" + query + "%" + " OR PhoneNumber ILIKE" + "%" + query + "%" + " OR Id ILIKE" + "%" + query + "%" + ""
	db = db.Where(finalQuery)
	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Sayfa numarası geçersiz"})

	}

	perPage, err := strconv.Atoi(c.Query("per_page", "10"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Sayfa başına alınacak kayıt sayısı geçersiz"})

	}

	offset := (page - 1) * perPage
	limit := perPage
	db = db.Offset(offset).Limit(limit)

	if err := db.Find(&customer).Error; err != nil {
		return err

	}
	return c.JSON(customer)
}

// ShowCustomer godoc
// @Summary Show Customer
// @Description Paginate PaginateCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} api_structure.Customer
// @Router /test/customer/paginate
func (controller *CustomerController) PaginateCustomer(c *fiber.Ctx) error {

	return nil
}

// ShowCustomer godoc
// @Summary Show Customer
// @Description Infinite InfiniteCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} api_structure.Customer
// @Router /test/customer/infinite
func (controller *CustomerController) InfiniteCustomer(c *fiber.Ctx) error {

	return nil
}

// ShowCustomer godoc
// @Summary Show Customer
// @Description Create CreateCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} api_structure.Customer
// @Router /test/customer/create
func (controller *CustomerController) CreateCustomer(c *fiber.Ctx) error {
	data := api_structure.Customer{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateCustomer(data)

	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowCustomer godoc
// @Summary Show Customer
// @Description CreateBulk CreateBulkCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} api_structure.Customer
// @Router /test/customer/createbulk
func (controller *CustomerController) CreateBulkCustomer(c *fiber.Ctx) error {
	data := []api_structure.Customer{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateBulkCustomer(data)

	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowCustomer godoc
// @Summary Show Customer
// @Description Update UpdateCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} api_structure.Customer
// @Router /test/customer/update
func (controller *CustomerController) UpdateCustomer(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	editData := api_structure.Customer{}

	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateCustomer(id, editData)

	if uerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Update Data",
			"message": uerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"succes":  true,
		"message": "Updated Successfully",
		"type":    "Update Data",
	})
}

// ShowCustomer godoc
// @Summary Show Customer
// @Description Delete DeleteCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} api_structure.Customer
// @Router /test/customer/delete
func (controller *CustomerController) DeleteCustomer(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := controller.Svc.DeleteCustomer(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Delete Data",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted Successfully",
		"type":    "Delete Data",
		"success": true,
	})
}

// ShowCustomer godoc
// @Summary Show Customer
// @Description DeleteBulk DeleteBulkCustomer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Success 200 {object} Customer
// @Router /test/customer/deletebulk
func (controller *CustomerController) DeleteBulkCustomer(c *fiber.Ctx) error {

	return nil
}
