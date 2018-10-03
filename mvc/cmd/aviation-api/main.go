package main

import (
	"net/http"

	"github.com/belfastgophers/structure/mvc/api/rest"
	"github.com/belfastgophers/structure/mvc/repo/inmemory"
	"github.com/belfastgophers/structure/mvc/service"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	s := &http.Server{
		Addr: ":2241",
	}
	g := e.Group("/aviation/api")

	// repo setup
	airplaneRepo := inmemory.NewAirplaneStore()
	airportRepo := inmemory.NewAirportStore()

	// services setup
	airplaneSvc := service.NewAirplaneService(airplaneRepo)
	airportSvc := service.NewAirportService(airportRepo)

	// handlers setup
	airplaneHnd := rest.NewAirplaneHandler(airplaneSvc)
	airportHnd := rest.NewAirportHandler(airportSvc)

	// setup routes
	airplaneHnd.InitRoutes(g)
	airportHnd.InitRoutes(g)

	// startup server
	e.Logger.Fatal(e.StartServer(s))
}
