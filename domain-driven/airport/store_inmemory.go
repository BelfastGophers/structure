package airport

import (
	"context"
	"strings"

	"github.com/belfastgophers/structure/domain-driven"
	"github.com/pkg/errors"
)

// airportStore will get and save airports.
type airportStore struct {
	aa map[string]*aviation.Airport
}

// NewAirportStore will create and return a new airport store instance.
func NewAirportStore() *airportStore {
	return &airportStore{
		aa: getDefaultAirports(),
	}
}

// Airport gets a single airport from the data store.
func (r *airportStore) Airport(ctx context.Context, icao string) (*aviation.Airport, error) {
	if a, ok := r.aa[strings.ToLower(icao)]; ok {
		return a, nil
	}
	return nil, errors.Errorf("unable to find airport with ICAO %s", icao)
}

// Airports returns all stored airports.
func (r *airportStore) Airports(ctx context.Context) ([]aviation.Airport, error) {
	aa := make([]aviation.Airport, 0)
	for _, a := range r.aa {
		aa = append(aa, *a)
	}
	if len(aa) == 0 {
		return nil, errors.Errorf("unable to find airports")
	}
	return aa, nil
}

// AddAirport will add a new Airport to the datastore.
func (r *airportStore) AddAirport(ctx context.Context, airport aviation.Airport) (*aviation.Airport, error) {
	if _, ok := r.aa[strings.ToLower(airport.ICAO)]; ok {
		return nil, errors.Errorf("airplane already exists")
	}
	r.aa[strings.ToLower(airport.ICAO)] = &airport
	return &airport, nil
}

func getDefaultAirports() map[string]*aviation.Airport {
	return map[string]*aviation.Airport{
		"egad": {
			ICAO:     "EGAD",
			Name:     "Newtownards Airfield",
			Location: "Newtownards, Northern Ireland",
		},
	}
}
