package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Server
	e.GET("/api/players/:id", GetPlayer)
	e.GET("/health", Health)
	e.Logger.Fatal(e.Start(":9999"))

}

func Health(c echo.Context) error {
	return c.JSON(200, &HealthData{Status: "UP"})
}

type HealthData struct {
	Status string `json:"status,omitempty"`
}

func GetPlayer(c echo.Context) error {
	p := &Player{
		Name:  "Joe Doe",
		Email: "joe@doe.com",
	}
	return c.JSON(http.StatusOK, p)
}

type Player struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
