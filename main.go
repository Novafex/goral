package main

import (
	"fmt"
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
	data := Data{
		Name:        "Customer",
		Description: "Holds information about an individual customer",
		Definition: Definition{
			Fields: map[string]Field{
				"id": {
					Name:        "Id",
					Type:        "int",
					Description: "Identifier for the object",
					Optional:    false,
				},
				"name": {
					Name:        "Name",
					Type:        "string",
					Description: "Name of the customer",
					Optional:    false,
				},
				"phoneNumber": {
					Name:        "PhoneNumber",
					Type:        "string",
					Description: "Contact phone number",
					Optional:    true,
				},
			},
		},
		Search:      []string{"name", "phoneNumber"},
		Permissions: struct{ Key string }{Key: "real-estate:customer"},
		Actions: []string{
			"Get",
			"Paginate",
			"Infinite",
			"Create",
			"Create_bulk",
			"Update",
			"Delete",
			"Delete_bulk",
		},
	}

	// Klasörleri oluştur
	createDirectories([]string{"goral"})

	generateStructFile(data)
	generateServiceFile(data)
	generateControllerFile(data)

	cmd := exec.Command("gofmt", "-w", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}

func createDirectories(directories []string) {
	for _, dir := range directories {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Klasör oluşturma hatası: %s - %s\n", dir, err)
		}
	}
}

func generateStructFile(data Data) {
	fileName := "goral/" + strings.ToLower(data.Name) + "_types.go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Dosya oluşturma hatası:", err)
		return
	}
	defer file.Close()

	file.WriteString("package goral\n\n")
	file.WriteString(fmt.Sprintf("type %s struct {\n", data.Name))
	for fieldName, fieldData := range data.Definition.Fields {
		jsonTag := fieldName
		if fieldData.Optional {
			jsonTag += ",omitempty"
		}
		file.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\t// %s\n", fieldData.Name, fieldData.Type, jsonTag, fieldData.Description))
	}
	file.WriteString("}\n")
	fmt.Printf("%s dosyası başarıyla oluşturuldu.\n", fileName)
}

func generateServiceFile(data Data) {
	fileName := "goral/" + strings.ToLower(data.Name) + "_services.go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Dosya oluşturma hatası:", err)
		return
	}
	defer file.Close()

	file.WriteString("package goral\n\n")
	file.WriteString(`import "gorm.io/gorm"`)
	file.WriteString(fmt.Sprintf("\n\n type I%s interface {\n", data.Name))

	for _, action := range data.Actions {
		funcName := strings.ToLower(action) + data.Name

		switch action {
		case "Get":
			file.WriteString(fmt.Sprintf("%s(filter %s) ([]%s, error) \n", funcName, data.Name, data.Name))
		case "Paginate":
			file.WriteString(fmt.Sprintf("%s(filter %s) ([]%s, error) \n", funcName, data.Name, data.Name))
		case "Infinite":
			file.WriteString(fmt.Sprintf("%s(filter %s) ([]%s, error) \n", funcName, data.Name, data.Name))
		case "Create":
			file.WriteString(fmt.Sprintf("%s(data %s) (%s, error) \n", funcName, data.Name, data.Name))
		case "Create_bulk":
			file.WriteString(fmt.Sprintf("%s(data []%s) ([]%s, error) \n", funcName, data.Name, data.Name))
		case "Update":
			file.WriteString(fmt.Sprintf("%s(id int, data %s) error\n", funcName, data.Name))
		case "Delete":
			file.WriteString(fmt.Sprintf("%s(id int) error \n", funcName))
		case "Delete_bulk":
			file.WriteString(fmt.Sprintf("%s(id int) error \n", funcName))
		}
	}

	file.WriteString("}\n")

	file.WriteString(fmt.Sprintf("type %sService struct{ DB *gorm.DB }\n", data.Name))

	for _, action := range data.Actions {
		funcName := strings.ToLower(action) + data.Name
		file.WriteString(fmt.Sprintf("func (c *%sService) %s", data.Name, funcName))
		switch action {
		case "Get":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) {return []%s{},nil \n}\n\n", data.Name, data.Name, data.Name))
		case "Paginate":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) {return []%s{},nil \n}\n\n", data.Name, data.Name, data.Name))
		case "Infinite":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) {return []%s{},nil \n}\n\n", data.Name, data.Name, data.Name))
		case "Create":
			file.WriteString(fmt.Sprintf("(data %s) (%s, error) {return %s{},nil \n}\n\n", data.Name, data.Name, data.Name))
		case "Create_bulk":
			file.WriteString(fmt.Sprintf("(data []%s) ([]%s, error) {return []%s{},nil \n}\n\n", data.Name, data.Name, data.Name))
		case "Update":
			file.WriteString(fmt.Sprintf("(id int, data %s) error {return nil\n}\n\n", data.Name))
		case "Delete":
			file.WriteString("(id int) error {return nil\n}\n\n")
		case "Delete_bulk":
			file.WriteString("(id int) error {return nil\n}\n\n")
		}
	}

	fmt.Printf("%s dosyası başarıyla oluşturuldu.\n", fileName)
}

func generateControllerFile(data Data) {
	fileName := "goral/" + strings.ToLower(data.Name) + "_endpoints.go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Dosya oluşturma hatası:", err)
		return
	}
	defer file.Close()

	file.WriteString("package goral\n\n")
	file.WriteString(`import "github.com/gofiber/fiber/v2"`)

	file.WriteString(fmt.Sprintf("\n\n type %sController struct{ Svc %sService }\n", data.Name, data.Name))

	for _, action := range data.Actions {
		functionName := action + data.Name
		file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error {return nil}\n", data.Name, functionName))
	}

	fmt.Printf("%s dosyası başarıyla oluşturuldu.\n", fileName)
}
