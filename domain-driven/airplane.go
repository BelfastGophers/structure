package aviation

import "context"

// Airplane is a single airplane entity.
type Airplane struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	Registration string `json:"registration"`
	YearBuilt    int    `json:"yearBuilt"`
	Owner        string `json:"owner"`
}

// AirplaneService can be passed to transports and implemented by a service.
type AirplaneService interface {
	AirplaneRepository
}

type AirplaneRepository interface {
	Airplane(ctx context.Context, reg string) (*Airplane, error)
	Airplanes(ctx context.Context) ([]Airplane, error)
	AddAirplane(ctx context.Context, airplane Airplane) (*Airplane, error)
}
