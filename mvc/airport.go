package aviation

import "context"

//  Airport defines a single airport that can be stored.
type Airport struct {
	ICAO     string
	Name     string
	Location string
}

// AirportService defines a service used to access airports.
// This model could be different to the repo interface depending on requirements.
type AirportService interface {
	AirportRepository
}

// AirportRepository defines a data store for an Airport.
type AirportRepository interface {
	Airport(ctx context.Context, icao string) (*Airport, error)
	Airports(ctx context.Context) ([]Airport, error)
	AddAirport(ctx context.Context, airport Airport) (*Airport, error)
}
