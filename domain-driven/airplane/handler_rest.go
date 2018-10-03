package airplane

import (
	"net/http"

	"github.com/belfastgophers/structure/domain-driven"

	"github.com/labstack/echo"
)

const (
	airplanesURL = "/airplanes"
	airplaneURL  = "/airplanes/:registration"
)

type airplaneHandler struct {
	svc aviation.AirplaneService
}

// NewAirplaneHandler creates a new airplaneHandler.
func NewAirplaneHandler(svc aviation.AirplaneService) *airplaneHandler {
	return &airplaneHandler{
		svc: svc,
	}
}

// InitRoutes will setup routes on the aviation api group.
func (h *airplaneHandler) InitRoutes(group *echo.Group) {
	group.GET(airplaneURL, h.airplane)
	group.GET(airplanesURL, h.airplanes)
	group.POST(airplanesURL, h.add)
}

func (h *airplaneHandler) airplane(c echo.Context) error {
	a, err := h.svc.Airplane(c.Request().Context(), c.Param("registration"))
	if err != nil {
		return &echo.HTTPError{
			Message:  err.Error(),
			Code:     http.StatusInternalServerError,
			Internal: err,
		}
	}
	return c.JSON(http.StatusOK, a)
}

func (h *airplaneHandler) airplanes(c echo.Context) error {
	a, err := h.svc.Airplanes(c.Request().Context())
	if err != nil {
		return &echo.HTTPError{
			Message:  err.Error(),
			Code:     http.StatusInternalServerError,
			Internal: err,
		}
	}
	return c.JSON(http.StatusOK, a)
}

func (h *airplaneHandler) add(c echo.Context) error {
	var req aviation.Airplane
	if err := c.Bind(&req); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "invalid reequest object",
			Internal: err,
		}
	}
	airplane, err := h.svc.AddAirplane(c.Request().Context(), req)
	if err != nil {
		return &echo.HTTPError{
			Message:  err.Error(),
			Code:     http.StatusInternalServerError,
			Internal: err,
		}
	}
	return c.JSON(http.StatusCreated, airplane)
}
