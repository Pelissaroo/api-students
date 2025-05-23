package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("students/:id", deleteStudent)

	// 	Routes:
	// - GET /students - List all students
	// - POST /students - Create student
	// - GET /students/:id - Get infos from a specifc student
	// - PUT /students/:id - Update student
	// - DELETE /students/:id - Delete student

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
	//USAR CURL.EXE NO TERMINAL
}

// Handler
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "list of all students")
}

func createStudent(c echo.Context) error {
	return c.String(http.StatusOK, "Creat student")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Get %s Student", id)
	return c.String(http.StatusOK, getStud)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update %s Student", id)
	return c.String(http.StatusOK, updateStud)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delete %s Student", id)
	return c.String(http.StatusOK, deleteStud)
}
