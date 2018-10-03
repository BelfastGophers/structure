package airport

import (
	"net/http"

	"github.com/belfastgophers/structure/domain-driven"

	"github.com/labstack/echo"
)

const (
	airportsURL = "/airports"
	airportURL  = "/airports/:icao"
)

type airportHandler struct {
	svc aviation.AirportService
}

// NewAirportHandler will return a new airportHandler
func NewAirportHandler(svc aviation.AirportService) *airportHandler {
	return &airportHandler{
		svc: svc,
	}
}

// InitRoutes will setup routes on the aviation api group.
func (h *airportHandler) InitRoutes(group *echo.Group) {
	group.GET(airportURL, h.airport)
	group.GET(airportsURL, h.airports)
	group.POST(airportsURL, h.add)
}

func (h *airportHandler) airport(c echo.Context) error {
	a, err := h.svc.Airport(c.Request().Context(), c.Param("icao"))
	if err != nil {
		return &echo.HTTPError{
			Message:  err.Error(),
			Code:     http.StatusInternalServerError,
			Internal: err,
		}
	}
	return c.JSON(http.StatusOK, a)
}

func (h *airportHandler) airports(c echo.Context) error {
	a, err := h.svc.Airports(c.Request().Context())
	if err != nil {
		return &echo.HTTPError{
			Message:  err.Error(),
			Code:     http.StatusInternalServerError,
			Internal: err,
		}
	}
	return c.JSON(http.StatusOK, a)
}

func (h *airportHandler) add(c echo.Context) error {
	var req aviation.Airport
	if err := c.Bind(&req); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "invalid request object",
			Internal: err,
		}
	}
	airplane, err := h.svc.AddAirport(c.Request().Context(), req)
	if err != nil {
		return &echo.HTTPError{
			Message:  err.Error(),
			Code:     http.StatusInternalServerError,
			Internal: err,
		}
	}
	return c.JSON(http.StatusCreated, airplane)
}
