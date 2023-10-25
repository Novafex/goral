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
	var data Data

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)

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
		if !fieldData.Relational {
			jsonTag := fieldName

			if fieldData.Optional {
				jsonTag += ",omitempty"
			}

			gormTag := fmt.Sprintf("gorm:\"column:%s;type:%s\"", fieldData.Name, fieldData.Type)
			jsonTag = fmt.Sprintf("json:\"%s\"", fieldName)
			swaggertypeTag := "swaggertype:\"string\""

			file.WriteString(fmt.Sprintf("\t%s %s `%s %s %s`\t// %s\n", fieldData.Name, fieldData.Type, gormTag, jsonTag, swaggertypeTag, fieldData.Description))
		} else {
			jsonTag := fieldName
			gormTag := fmt.Sprintf("gorm:\"column:%s;type:%s\"", fieldData.Name, "nullable")
			jsonTag = fmt.Sprintf("json:\"%s\"", fieldName)
			swaggertypeTag := "swaggertype:\"string\""

			file.WriteString(fmt.Sprintf("\t%s %s `%s %s %s`\t// %s\n", fieldData.Name, "*"+fieldData.Description, gormTag, jsonTag, swaggertypeTag, fieldData.Description))
		}
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
	file.WriteString(`import (
        "gorm.io/gorm" 
		"gorm.io/gorm/clause"
    )`)
	file.WriteString(fmt.Sprintf("\n\n type I%s interface {\n", data.Name))

	for _, action := range data.Actions {
		funcName := action + data.Name

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
		funcName := action + data.Name
		file.WriteString(fmt.Sprintf("func (c *%sService) %s", data.Name, funcName))
		switch action {
		case "Get":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) { \n", data.Name, data.Name))
			file.WriteString(fmt.Sprintf("result := []%s{} \n", data.Name))
			file.WriteString("var err error \n")
			file.WriteString(fmt.Sprintf("if err = c.DB.Preload(clause.Associations).Model(&%s{}).Where(filter).Find(&result).Error; err != nil { \nreturn result, err \n}", data.Name))
			file.WriteString("\nreturn result, err} \n\n")
		case "Paginate":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) {return []%s{},nil \n}\n\n", data.Name, data.Name, data.Name))
		case "Infinite":
			file.WriteString(fmt.Sprintf("(filter %s) ([]%s, error) {return []%s{},nil \n}\n\n", data.Name, data.Name, data.Name))
		case "Create":
			file.WriteString(fmt.Sprintf("(data %s) (%s, error) { \n", data.Name, data.Name))
			file.WriteString("var err error \n")
			file.WriteString("if err =  c.DB.Create(&data).Error; err != nil { \nreturn data, err \n}")
			file.WriteString("\nreturn data, err} \n\n")
		case "CreateBulk":
			file.WriteString(fmt.Sprintf("(data []%s) ([]%s, error) { \n", data.Name, data.Name))
			file.WriteString("var err error \n")
			file.WriteString("if err = c.DB.CreateInBatches(&data, len(data)).Error; err != nil { \nreturn data, err \n}")
			file.WriteString("\nreturn data, err} \n\n")
		case "Update":
			file.WriteString(fmt.Sprintf("(id int, data %s) error {\n\n", data.Name))
			file.WriteString("var err error \n")
			file.WriteString(fmt.Sprintf("if err = c.DB.Model(%s{}).Where(\"id = ?\", id).Updates(&data).Error; err != nil {\nreturn err\n}", data.Name))
			file.WriteString("\nreturn err} \n\n")
		case "Delete":
			file.WriteString("(id int) error { \n")
			file.WriteString("var err error \n")
			file.WriteString(fmt.Sprintf("if err = c.DB.Where(\"id = ?\", id).Delete(&%s{}).Error; err != nil { \nreturn err \n}", data.Name))
			file.WriteString("\nreturn err} \n\n")
		case "DeleteBulk":
			file.WriteString("(ids []int) error {\n")
			file.WriteString("var err error \n")
			file.WriteString("for _, id := range ids {\n")
			file.WriteString(fmt.Sprintf("if err = c.DB.Where(\"id = ?\", id).Delete(&%s{}).Error; err != nil {\nreturn err\n}\n}", data.Name))
			file.WriteString("\nreturn err} \n\n")
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
	file.WriteString(`import (
        "github.com/gofiber/fiber/v2"
    )`)

	file.WriteString(fmt.Sprintf("\n\n type %sController struct{ Svc %sService }\n", data.Name, data.Name))

	for _, action := range data.Actions {
		functionName := action + data.Name
		file.WriteString(fmt.Sprintf("// Show%s godoc\n", data.Name))
		file.WriteString(fmt.Sprintf("// @Summary Show %s\n", data.Name))
		file.WriteString(fmt.Sprintf("// @Description %s %s\n", action, functionName))
		file.WriteString(fmt.Sprintf("// @Tags %s\n", data.Name))
		file.WriteString(fmt.Sprintf("// @Param id path string true \"%s ID\"\n", data.Name))
		file.WriteString(fmt.Sprintf("// @Success 200 {object} %s\n", data.Name))
		file.WriteString(fmt.Sprintf("// @Router /test/%s/%s\n", strings.ToLower(data.Name), strings.ToLower(action)))
		file.WriteString(fmt.Sprintf("func (controller *%sController) %s(c *fiber.Ctx) error {return nil}\n", data.Name, functionName))
	}

	fmt.Printf("%s dosyası başarıyla oluşturuldu.\n", fileName)
}
