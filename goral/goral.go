package goral

import (
	api_controller "generate/goral/controllers"
	api_service "generate/goral/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Run(app fiber.Router, DB *gorm.DB) {

	Customer := api_controller.CustomerController{Svc: api_service.CustomerService{DB: DB}}
	app_Customer := app.Group("/customer")
	app_Customer.Get("/get", Customer.GetCustomer)
	app_Customer.Get("/paginate", Customer.PaginateCustomer)
	app_Customer.Get("/infinite", Customer.InfiniteCustomer)
	app_Customer.Post("/create", Customer.CreateCustomer)
	app_Customer.Post("/createbulk", Customer.CreateBulkCustomer)
	app_Customer.Put("/update", Customer.UpdateCustomer)
	app_Customer.Delete("/delete", Customer.DeleteCustomer)
	app_Customer.Delete("/deletebulk", Customer.DeleteBulkCustomer)

	Advisor := api_controller.AdvisorController{Svc: api_service.AdvisorService{DB: DB}}
	app_Advisor := app.Group("/advisor")
	app_Advisor.Get("/get", Advisor.GetAdvisor)
	app_Advisor.Get("/paginate", Advisor.PaginateAdvisor)
	app_Advisor.Get("/infinite", Advisor.InfiniteAdvisor)
	app_Advisor.Post("/create", Advisor.CreateAdvisor)
	app_Advisor.Post("/createbulk", Advisor.CreateBulkAdvisor)
	app_Advisor.Put("/update", Advisor.UpdateAdvisor)
	app_Advisor.Delete("/delete", Advisor.DeleteAdvisor)
	app_Advisor.Delete("/deletebulk", Advisor.DeleteBulkAdvisor)

}
