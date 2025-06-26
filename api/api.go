package api

import (
	"github.com/Pelissaroo/api-students/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/Pelissaroo/api-students/docs"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

// @title Student API
// @version 1.0
// @description This is a sample server Student API
// @host localhost:8080
// @BasePath /
// @schemes http
func NewServer() *API {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.GetStudents)
	api.Echo.POST("/students", api.CreateStudent)
	api.Echo.GET("/students/:id", api.GetStudent)
	api.Echo.PUT("/students/:id", api.UpdateStudent)
	api.Echo.DELETE("students/:id", api.DeleteStudent)
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
