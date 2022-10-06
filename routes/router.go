package routes

import (
	"project-1/controllers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	// echoSwagger "github.com/swaggo/echo-swagger/v2"
	// "github.com/swaggo/echo-swagger/v2"
	// _ "github.com/swaggo/echo-swagger/v2/example/docs"
)

func New() *echo.Echo{
	e :=echo.New()
    e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/book",controllers.CreateBook)
	e.GET("/book",controllers.GetBookController)
	e.GET("/book/:id",controllers.DetailBookController)
	e.PUT("/book/:id",controllers.UpdateBook)
	e.DELETE("/book/:id",controllers.DeleteBook)
    return e


}
