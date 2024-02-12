package goral

import (
        "github.com/gofiber/fiber/v2"
		"strconv"
		api_service "generate/goral/services"
		api_structure "generate/goral/structures"
    )

 type AdvisorController struct{ Svc api_service.AdvisorService }
// ShowAdvisor godoc
// @Summary Show Advisor
// @Description Get GetAdvisor
// @Tags Advisor
// @Param id path string true "Advisor ID"
// @Success 200 {object} api_structure.Advisor
// @Router  /advisor/get   [GET]
func (controller *AdvisorController) GetAdvisor(c *fiber.Ctx) error { 
var advisor []api_structure.Advisor

db := controller.Svc.DB.Table("Advisor")
query := c.Query("query")
finalQuery :="Id ILIKE"+"%"+ query +"%"+" OR Name ILIKE"+"%"+ query +"%"+" OR PhoneNumber ILIKE"+"%"+ query +"%"+"" 
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

if err := db.Find(&advisor).Error; err != nil {return err

}
return c.JSON(advisor)
}

// ShowAdvisor godoc
// @Summary Show Advisor
// @Description Paginate PaginateAdvisor
// @Tags Advisor
// @Param id path string true "Advisor ID"
// @Success 200 {object} api_structure.Advisor
// @Router  /advisor/paginate  [GET]
func (controller *AdvisorController) PaginateAdvisor(c *fiber.Ctx) error { 

return nil
}

// ShowAdvisor godoc
// @Summary Show Advisor
// @Description Infinite InfiniteAdvisor
// @Tags Advisor
// @Param id path string true "Advisor ID"
// @Success 200 {object} api_structure.Advisor
// @Router  /advisor/infinite  [GET]
func (controller *AdvisorController) InfiniteAdvisor(c *fiber.Ctx) error { 

return nil
}

// ShowAdvisor godoc
// @Summary Show Advisor
// @Description Create CreateAdvisor
// @Tags Advisor
// @Param id path string true "Advisor ID"
// @Success 200 {object} api_structure.Advisor
// @Router  /advisor/create  [POST]
func (controller *AdvisorController) CreateAdvisor(c *fiber.Ctx) error { 
data := api_structure.Advisor{}

if err := c.BodyParser(&data); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
"type":    "Invalid Data",
"message": err.Error(),
})
}

result, rerr := controller.Svc.CreateAdvisor(data)

			if rerr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"type":    "Create Data",
					"message": rerr.Error(),
				})
			}

return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAdvisor godoc
// @Summary Show Advisor
// @Description CreateBulk CreateBulkAdvisor
// @Tags Advisor
// @Param id path string true "Advisor ID"
// @Success 200 {object} api_structure.Advisor
// @Router  /advisor/createbulk   [POST]
func (controller *AdvisorController) CreateBulkAdvisor(c *fiber.Ctx) error { 
data := []api_structure.Advisor{}

if err := c.BodyParser(&data); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
"type":    "Invalid Data",
"message": err.Error(),
})
}

result, rerr := controller.Svc.CreateBulkAdvisor(data)

			if rerr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"type":    "Create Data",
					"message": rerr.Error(),
				})
			}

return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAdvisor godoc
// @Summary Show Advisor
// @Description Update UpdateAdvisor
// @Tags Advisor
// @Param id path string true "Advisor ID"
// @Success 200 {object} api_structure.Advisor
// @Router  /advisor/update  [PUT]
func (controller *AdvisorController) UpdateAdvisor(c *fiber.Ctx) error { 
id, _ := strconv.Atoi(c.Params("id"))

editData:= api_structure.Advisor{}

if err := c.BodyParser(&editData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"type": "Invalid Data",
				"message": err.Error(),
			})}

uerr := controller.Svc.UpdateAdvisor(id, editData)

if uerr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Update Data",
				"message": uerr.Error(),
			})}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"succes":  true,
				"message": "Updated Successfully",
				"type":    "Update Data",
			})}

// ShowAdvisor godoc
// @Summary Show Advisor
// @Description Delete DeleteAdvisor
// @Tags Advisor
// @Param id path string true "Advisor ID"
// @Success 200 {object} api_structure.Advisor
// @Router  /advisor/delete   [DELETE]
func (controller *AdvisorController) DeleteAdvisor(c *fiber.Ctx) error { 
id, _ := strconv.Atoi(c.Params("id"))



err := controller.Svc.DeleteAdvisor(id)

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

// ShowAdvisor godoc
// @Summary Show Advisor
// @Description DeleteBulk DeleteBulkAdvisor
// @Tags Advisor
// @Param id path string true "Advisor ID"
// @Success 200 {object} Advisor
// @Router  /advisor/deletebulk  [DELETE]
func (controller *AdvisorController) DeleteBulkAdvisor(c *fiber.Ctx) error { 

return nil
}