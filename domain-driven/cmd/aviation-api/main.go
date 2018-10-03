package main

import (
	"net/http"

	"github.com/belfastgophers/structure/domain-driven/airplane"
	"github.com/belfastgophers/structure/domain-driven/airport"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	s := &http.Server{
		Addr: ":2241",
	}
	g := e.Group("/aviation/api")

	// airplane setup
	airplaneRepo := airplane.NewAirplaneStore()
	airplaneSvc := airplane.NewAirplaneService(airplaneRepo)
	airplaneHnd := airplane.NewAirplaneHandler(airplaneSvc)
	airplaneHnd.InitRoutes(g)

	// airport setup
	airportRepo := airport.NewAirportStore()
	airportSvc := airport.NewAirportService(airportRepo)
	airportHnd := airport.NewAirportHandler(airportSvc)
	airportHnd.InitRoutes(g)

	// startup server
	e.Logger.Fatal(e.StartServer(s))
}
