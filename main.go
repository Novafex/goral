package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Field struct {
	Name        string
	Type        string
	Description string
	Optional    bool
	Relational  bool
}

type Definition struct {
	Fields map[string]Field
}

type ServiceAction struct {
	Name string
}

type ServiceDefinition struct {
	Actions []ServiceAction
}

type Data struct {
	Name        string
	Description string
	Definition  Definition
	Search      []string
	Permissions struct {
		Key string
	}
	Actions []string
}

func main() {

	var first string
	fmt.Println("Process: \n1- data \n2- integration")
	fmt.Scanln(&first)
	if first == "1" {
		var datas []Data
		jsonFile, err := os.Open("data.json")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully Opened data.json")
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &datas)

		// Klasörleri oluştur
		createDirectories([]string{"goral", "goral/controllers", "goral/services", "goral/structures"})
		for _, data := range datas {
			generateStructFile(data)
			generateServiceFile(data)
			generateControllerFile(data)
		}
		generateRouterFile(datas)
		// gofmt komutunu çalıştır
		gofmtCmd := exec.Command("gofmt", "-w", ".")
		if err := gofmtCmd.Run(); err != nil {
			log.Fatal(err)
		}
		// swag init komutunu çalıştır
		swagCmd := exec.Command("swag", "init")
		if err := swagCmd.Run(); err != nil {
			log.Fatal(err)
		}
		// go generate komutunu çalıştır
		generateCmd := exec.Command("go", "generate", "./goral/services")
		if err := generateCmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else if first == "2" {
		log.Println("Integration being developed.")
	} else {
		log.Println("WARNING!!!")
	}
}

func createDirectories(directories []string) {
	for _, dir := range directories {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Klasör oluşturma hatası: %s - %s\n", dir, err)
		}
	}
}

func createFieldTags(fieldName, columnName, typeName string) (string, string) {
	jsonTag := fmt.Sprintf("json:\"%s\"", fieldName)
	gormTag := fmt.Sprintf("gorm:\"column:%s;type:%s\"", columnName, typeName)
	return jsonTag, gormTag
}

func generateStructFile(data Data) {
	filePath := "goral/structures/"
	fileName := filePath + strings.ToLower(data.Name) + "_types.go"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Dosya oluşturma hatası: %v", err)
		return
	}
	defer file.Close()

	structName := strings.Title(data.Name)
	file.WriteString("package goral\n\n")
	file.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	for fieldName, fieldData := range data.Definition.Fields {
		jsonTag, gormTag := createFieldTags(fieldName, fieldData.Name, fieldData.Type)

		if !fieldData.Relational {
			swaggertypeTag := "swaggertype:\"string\""
			file.WriteString(fmt.Sprintf("\t%s %s `%s %s %s`\t// %s\n", fieldData.Name, fieldData.Type, gormTag, jsonTag, swaggertypeTag, fieldData.Description))
		} else {
			jsonTag, gormTag = createFieldTags(fieldName, fieldData.Name, "nullable")
			swaggertypeTag := "swaggertype:\"string\""
			file.WriteString(fmt.Sprintf("\t%s *%s `%s %s %s`\t// %s\n", fieldData.Name, fieldData.Description, gormTag, jsonTag, swaggertypeTag, fieldData.Description))
		}
	}
	file.WriteString("}\n")
	fmt.Printf("%s dosyası başarıyla oluşturuldu.\n", fileName)
}

func generateServiceFile(data Data) {
	fileName := "goral/services/" + strings.ToLower(data.Name) + "_services.go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Dosya oluşturma hatası:", err)
		return
	}
	defer file.Close()

	file.WriteString("package goral\n\n")
	file.WriteString(`import (
        "gorm.io/gorm" 
		"gorm.io/gorm/clause"
		api_structure "generate/goral/structures"
    )`)
	file.WriteString(fmt.Sprintf("\n //go:generate mockgen -destination=../mocks/services/%s.go -package=goral generate/goral/services I%s\n", strings.ToLower(data.Name), data.Name))

	file.WriteString(fmt.Sprintf("type %sService struct{ DB *gorm.DB }\n", data.Name))
	file.WriteString(fmt.Sprintf("\n type I%s interface {\n", data.Name))

	for _, action := range data.Actions {
		funcName := action + data.Name

		switch action {
		case "Get":
			file.WriteString(fmt.Sprintf("%s(filter %s) ([]%s, error) \n", funcName, "api_structure."+data.Name, "api_structure."+data.Name))
		case "Paginate":
			file.WriteString(fmt.Sprintf("%s(filter %s) ([]%s, error) \n", funcName, "api_structure."+data.Name, "api_structure."+data.Name))
		case "Infinite":
			file.WriteString(fmt.Sprintf("%s(filter %s) ([]%s, error) \n", funcName, "api_structure."+data.Name, "api_structure."+data.Name))
		case "Create":
			file.WriteString(fmt.Sprintf("%s(data %s) (%s, error) \n", funcName, "api_structure."+data.Name, "api_structure."+data.Name))
		case "Create_bulk":
			file.WriteString(fmt.Sprintf("%s(data []%s) ([]%s, error) \n", funcName, "api_structure."+data.Name, "api_structure."+data.Name))
		case "Update":
			file.WriteString(fmt.Sprintf("%s(id int, data %s) error\n", funcName, "api_structure."+data.Name))
		case "Delete":
			file.WriteString(fmt.Sprintf("%s(id int) error \n", funcName))
		case "Delete_bulk":
			file.WriteString(fmt.Sprintf("%s(id int) error \n", funcName))
		}
	}

	file.WriteString("}\n")

	for _, action := range data.Actions {
		funcName := action + data.Name
		file.WriteString(fmt.Sprintf("func (c *%sService) %s", data.Name, funcName))
		switch action {
		case "Get":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) { \n", "api_structure."+data.Name, "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("result := []%s{} \n", "api_structure."+data.Name))
			file.WriteString("var err error \n")
			file.WriteString(fmt.Sprintf("if err = c.DB.Preload(clause.Associations).Model(&%s{}).Where(filter).Find(&result).Error; err != nil { \nreturn result, err \n}", "api_structure."+data.Name))
			file.WriteString("\nreturn result, err} \n\n")
		case "Paginate":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) {return []%s{},nil \n}\n\n", "api_structure."+data.Name, "api_structure."+data.Name, "api_structure."+data.Name))
		case "Infinite":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) {return []%s{},nil \n}\n\n", "api_structure."+data.Name, "api_structure."+data.Name, "api_structure."+data.Name))
		case "Create":
			file.WriteString(fmt.Sprintf("(data %s) (%s, error) { \n", "api_structure."+data.Name, "api_structure."+data.Name))
			file.WriteString("var err error \n")
			file.WriteString("if err =  c.DB.Create(&data).Error; err != nil { \nreturn data, err \n}")
			file.WriteString("\nreturn data, err} \n\n")
		case "CreateBulk":
			file.WriteString(fmt.Sprintf("(data []%s) ([]%s, error) { \n", "api_structure."+data.Name, "api_structure."+data.Name))
			file.WriteString("var err error \n")
			file.WriteString("if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil { \nreturn data, err \n}")
			file.WriteString("\nreturn data, err} \n\n")
		case "Update":
			file.WriteString(fmt.Sprintf("(id int, data %s) error {\n\n", "api_structure."+data.Name))
			file.WriteString("var err error \n")
			file.WriteString(fmt.Sprintf("if err = c.DB.Model(%s{}).Where(\"id = ?\", id).Updates(&data).Error; err != nil {\nreturn err\n}", "api_structure."+data.Name))
			file.WriteString("\nreturn err} \n\n")
		case "Delete":
			file.WriteString("(id int) error { \n")
			file.WriteString("var err error \n")
			file.WriteString(fmt.Sprintf("if err = c.DB.Where(\"id = ?\", id).Delete(&%s{}).Error; err != nil { \nreturn err \n}", "api_structure."+data.Name))
			file.WriteString("\nreturn err} \n\n")
		case "DeleteBulk":
			file.WriteString("(ids []int) error {\n")
			file.WriteString("var err error \n")
			file.WriteString("for _, id := range ids {\n")
			file.WriteString(fmt.Sprintf("if err = c.DB.Where(\"id = ?\", id).Delete(&%s{}).Error; err != nil {\nreturn err\n}\n}", "api_structure."+data.Name))
			file.WriteString("\nreturn err} \n\n")
		}
	}

	fmt.Printf("%s dosyası başarıyla oluşturuldu.\n", fileName)
}

func generateControllerFile(data Data) {
	fileName := "goral/controllers/" + strings.ToLower(data.Name) + "_endpoints.go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Dosya oluşturma hatası:", err)
		return
	}
	defer file.Close()

	file.WriteString("package goral\n\n")
	file.WriteString(`import (
        "github.com/gofiber/fiber/v2"
		"strconv"
		api_service "generate/goral/services"
		api_structure "generate/goral/structures"
    )`)

	file.WriteString(fmt.Sprintf("\n\n type %sController struct{ Svc %sService }\n", data.Name, "api_service."+data.Name))

	for _, action := range data.Actions {
		functionName := action + data.Name
		switch action {
		case "Get":
			file.WriteString(fmt.Sprintf("// Show%s godoc\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
			file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("// @Router  /%s/%s   [GET]\n", strings.ToLower(data.Name), strings.ToLower(action)))
			file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error { \n", data.Name, functionName))

			file.WriteString(fmt.Sprintf("var %s []%s\n\n", strings.ToLower(data.Name), "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("db := controller.Svc.DB.Table(\"%s\")\n", data.Name))
			file.WriteString("query := c.Query(\"query\")\n")
			var queryParts []string
			var finalQuery string
			for fieldName, fieldData := range data.Definition.Fields {
				if fieldData.Relational {
					file.WriteString(fmt.Sprintf("%s := c.Query(\"%s\")\n", fieldName, fieldName))
					file.WriteString(fmt.Sprintf("db = db.Where(\"%s = ?\", %s)\n\n", fieldName, fieldName))
				} else {

					queryPart := fieldData.Name + ` ILIKE"+"%"+ query +"%"+"`
					queryParts = append(queryParts, queryPart)
					finalQuery = strings.Join(queryParts, " OR ")

				}
			}
			file.WriteString(fmt.Sprintf("finalQuery :=\"%s\" \n", finalQuery))
			file.WriteString("db = db.Where(finalQuery)\n")
			file.WriteString("page, err := strconv.Atoi(c.Query(\"page\", \"1\"))\n\n")
			file.WriteString("if err != nil {\n\nreturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\n\"error\": \"Sayfa numarası geçersiz\"})\n\n}\n\n")

			file.WriteString("perPage, err := strconv.Atoi(c.Query(\"per_page\", \"10\"))\n\n")
			file.WriteString("if err != nil { \nreturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\"error\": \"Sayfa başına alınacak kayıt sayısı geçersiz\"})\n\n\n}")

			file.WriteString("\n\noffset := (page - 1) * perPage\nlimit := perPage\n")
			file.WriteString("db = db.Offset(offset).Limit(limit)\n\n")
			file.WriteString(fmt.Sprintf("if err := db.Find(&%s).Error; err != nil {return err\n\n}", strings.ToLower(data.Name)))
			file.WriteString(fmt.Sprintf("\nreturn c.JSON(%s)\n}", strings.ToLower(data.Name)))
		case "Paginate":
			file.WriteString(fmt.Sprintf("\n\n// Show%s godoc\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
			file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("// @Router  /%s/%s  [GET]\n", strings.ToLower(data.Name), strings.ToLower(action)))
			file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error { \n", data.Name, functionName))

			file.WriteString("\nreturn nil\n}")
		case "Infinite":
			file.WriteString(fmt.Sprintf("\n\n// Show%s godoc\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
			file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("// @Router  /%s/%s  [GET]\n", strings.ToLower(data.Name), strings.ToLower(action)))
			file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error { \n", data.Name, functionName))

			file.WriteString("\nreturn nil\n}")
		case "Create":
			file.WriteString(fmt.Sprintf("\n\n// Show%s godoc\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
			file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("// @Router  /%s/%s  [POST]\n", strings.ToLower(data.Name), strings.ToLower(action)))
			file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error { \n", data.Name, functionName))
			file.WriteString(fmt.Sprintf("data := %s{}\n\n", "api_structure."+data.Name))
			file.WriteString("if err := c.BodyParser(&data); err != nil {\nreturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\n\"type\":    \"Invalid Data\",\n\"message\": err.Error(),\n})\n}")
			file.WriteString(fmt.Sprintf("\n\nresult, rerr := controller.Svc.%s(data)\n", functionName))
			file.WriteString(`
			if rerr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"type":    "Create Data",
					"message": rerr.Error(),
				})
			}`)
			file.WriteString("\n\nreturn c.Status(fiber.StatusOK).JSON(result)\n}")
		case "CreateBulk":
			file.WriteString(fmt.Sprintf("\n\n// Show%s godoc\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
			file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("// @Router  /%s/%s   [POST]\n", strings.ToLower(data.Name), strings.ToLower(action)))
			file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error { \n", data.Name, functionName))
			file.WriteString(fmt.Sprintf("data := []%s{}\n\n", "api_structure."+data.Name))
			file.WriteString("if err := c.BodyParser(&data); err != nil {\nreturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\n\"type\":    \"Invalid Data\",\n\"message\": err.Error(),\n})\n}")
			file.WriteString(fmt.Sprintf("\n\nresult, rerr := controller.Svc.%s(data)\n", functionName))
			file.WriteString(`
			if rerr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"type":    "Create Data",
					"message": rerr.Error(),
				})
			}`)
			file.WriteString("\n\nreturn c.Status(fiber.StatusOK).JSON(result)\n}")
		case "Update":
			file.WriteString(fmt.Sprintf("\n\n// Show%s godoc\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
			file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("// @Router  /%s/%s  [PUT]\n", strings.ToLower(data.Name), strings.ToLower(action)))
			file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error { \n", data.Name, functionName))
			file.WriteString("id, _ := strconv.Atoi(c.Params(\"id\"))\n\n")
			file.WriteString(fmt.Sprintf("editData:= %s{}\n\n", "api_structure."+data.Name))
			file.WriteString(`if err := c.BodyParser(&editData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"type": "Invalid Data",
				"message": err.Error(),
			})}`)
			file.WriteString(fmt.Sprintf("\n\nuerr := controller.Svc.%s(id, editData)\n\n", functionName))
			file.WriteString(`if uerr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Update Data",
				"message": uerr.Error(),
			})}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"succes":  true,
				"message": "Updated Successfully",
				"type":    "Update Data",
			})}`)
		case "Delete":
			file.WriteString(fmt.Sprintf("\n\n// Show%s godoc\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
			file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", "api_structure."+data.Name))
			file.WriteString(fmt.Sprintf("// @Router  /%s/%s   [DELETE]\n", strings.ToLower(data.Name), strings.ToLower(action)))
			file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error { \n", data.Name, functionName))
			file.WriteString("id, _ := strconv.Atoi(c.Params(\"id\"))\n\n")
			file.WriteString(fmt.Sprintf("\n\nerr := controller.Svc.%s(id)\n\n", functionName))
			file.WriteString(`	if err != nil {
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
		}`)
		case "DeleteBulk":
			file.WriteString(fmt.Sprintf("\n\n// Show%s godoc\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
			file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", data.Name))
			file.WriteString(fmt.Sprintf("// @Router  /%s/%s  [DELETE]\n", strings.ToLower(data.Name), strings.ToLower(action)))
			file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error { \n", data.Name, functionName))

			file.WriteString("\nreturn nil\n}")
		}

	}

	fmt.Printf("%s dosyası başarıyla oluşturuldu.\n", fileName)
}

func generateRouterFile(datas []Data) {
	fileName := "goral/goral.go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Dosya oluşturma hatası:", err)
		return
	}
	defer file.Close()

	file.WriteString("package goral\n\n")
	file.WriteString(`import (
        "github.com/gofiber/fiber/v2"
		api_service "generate/goral/services"
		api_controller "generate/goral/controllers"
		"gorm.io/gorm"
    )`)
	file.WriteString(fmt.Sprintf("\nfunc Run(app fiber.Router, DB *gorm.DB,) { \n\n"))
	for _, data := range datas {
		file.WriteString(fmt.Sprintf("\n\n%s := api_controller.%sController{Svc: api_service.%sService{DB: DB}} \n", data.Name, data.Name, data.Name))
		appName := "app_" + data.Name
		file.WriteString(fmt.Sprintf("%s:=app.Group(\"/%s\")\n", appName, strings.ToLower(data.Name)))
		for _, action := range data.Actions {
			functionName := action + data.Name
			switch action {
			case "Get":
				file.WriteString(fmt.Sprintf("%s.Get(\"/%s\", %s.%s)\n", appName, strings.ToLower(action), data.Name, functionName))
			case "Paginate":
				file.WriteString(fmt.Sprintf("%s.Get(\"/%s\", %s.%s) \n", appName, strings.ToLower(action), data.Name, functionName))
			case "Infinite":
				file.WriteString(fmt.Sprintf("%s.Get(\"/%s\", %s.%s) \n", appName, strings.ToLower(action), data.Name, functionName))
			case "Create":
				file.WriteString(fmt.Sprintf("%s.Post(\"/%s\", %s.%s) \n", appName, strings.ToLower(action), data.Name, functionName))
			case "CreateBulk":
				file.WriteString(fmt.Sprintf("%s.Post(\"/%s\", %s.%s) \n", appName, strings.ToLower(action), data.Name, functionName))
			case "Update":
				file.WriteString(fmt.Sprintf("%s.Put(\"/%s\", %s.%s) \n", appName, strings.ToLower(action), data.Name, functionName))
			case "Delete":
				file.WriteString(fmt.Sprintf("%s.Delete(\"/%s\", %s.%s) \n", appName, strings.ToLower(action), data.Name, functionName))
			case "DeleteBulk":
				file.WriteString(fmt.Sprintf("%s.Delete(\"/%s\", %s.%s) \n", appName, strings.ToLower(action), data.Name, functionName))
			}

		}
	}
	file.WriteString("\n\n}")
}
