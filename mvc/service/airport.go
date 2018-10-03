package service

import (
	"context"
	"strings"

	"github.com/belfastgophers/structure/mvc"
	"github.com/pkg/errors"
)

type airportService struct {
	repo aviation.AirportRepository
}

func NewAirportService(repo aviation.AirportRepository) *airportService {
	return &airportService{
		repo: repo,
	}
}

// Airport will return an airport by ICAO if found.
// icao cannot be an empty string or > 6 chars.
func (s *airportService) Airport(ctx context.Context, icao string) (*aviation.Airport, error) {
	if err := s.validateICAO(icao); err != nil {
		return nil, errors.WithStack(err)
	}
	airplane, err := s.repo.Airport(ctx, icao)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return airplane, nil
}

// Airports will return all stored airports.
func (s *airportService) Airports(ctx context.Context) ([]aviation.Airport, error) {
	airports, err := s.repo.Airports(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return airports, nil
}

// AddAirport will add a new airport to the data store.
// airport.ICAO cannot be empty or > 6 chars.
func (s *airportService) AddAirport(ctx context.Context, airport aviation.Airport) (*aviation.Airport, error) {
	if err := s.validateICAO(airport.ICAO); err != nil {
		return nil, errors.WithStack(err)
	}
	retval, err := s.repo.AddAirport(ctx, airport)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return retval, nil
}

func (s *airportService) validateICAO(reg string) error {
	if strings.TrimSpace(reg) == "" {
		return errors.Errorf("empty registration submitted")
	}
	if len(strings.TrimSpace(reg)) > 6 {
		return errors.Errorf("invalid registration")
	}
	return nil
}
