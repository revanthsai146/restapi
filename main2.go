package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

// var Employees = []Employee{
// 	{ID: 2267841, Name: "Revanth", Salary: 75000},
// }

// GetOrders godoc
// @Summary Get  employee details
// @Description Get details of all employees
// @Tags employees
// @Accept  json
// @Produce  json
// @Success 200 {array} Employee
// @Router /employees [get]
func getEmployees(c echo.Context) error {
	var employee Employee
	employee.ID = 102
	employee.Name = "SAI"
	employee.Salary = 80000
	log.Print(employee)
	return c.JSON(http.StatusOK, employee)
}

// @title Employee API
// @version 1.0
// @description This is a sample service for employees
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()
	// Serve Swagger UI
	//e.Static("/", "docs")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// swaggerURL := echoSwagger.URL("/swagger.json")
	// e.GET("/swagger/*", echoSwagger.EchoWrapHandler(swaggerFiles.Handler, swaggerURL))

	// Generate and serve the Swagger UI HTML content
	// e.GET("/swagger-ui", func(c echo.Context) error {
	// 	swaggerConfig := echoSwagger.DefaultConfig()
	// 	swaggerConfig.URL = "/swagger.json"
	// 	htmlContent, err := swag.ReadDoc()
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, "Failed to generate Swagger UI")
	// 	}
	// 	return c.HTML(http.StatusOK, htmlContent)
	// })
	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, "Hello World")
	// })
	e.GET("/employees", getEmployees)
	e.Logger.Fatal(e.Start(":9031"))

}
