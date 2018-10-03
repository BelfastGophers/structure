package inmemory

import (
	"context"
	"strings"

	"github.com/belfastgophers/structure/mvc"
	"github.com/pkg/errors"
)

// airplaneStore contains methods to read and store airplanes.
type airplaneStore struct {
	aa map[string]*aviation.Airplane
}

// NewAirplaneStore will setup and return an inmemory airplaneStore
func NewAirplaneStore() *airplaneStore {
	return &airplaneStore{
		aa: getDefaultAirplanes(),
	}
}

// Airplane will return a single airplane by registration if found.
func (r *airplaneStore) Airplane(ctx context.Context, reg string) (*aviation.Airplane, error) {
	if a, ok := r.aa[strings.ToLower(reg)]; ok {
		return a, nil
	}
	return nil, errors.Errorf("unable to find aircraft with registration %s", reg)
}

// Airplanes will return all stored airplanes.
func (r *airplaneStore) Airplanes(ctx context.Context) ([]aviation.Airplane, error) {
	aa := make([]aviation.Airplane, 0)
	for _, a := range r.aa {
		aa = append(aa, *a)
	}
	return aa, nil
}

// AddAirplane will add a new airplane to the store.
func (r *airplaneStore) AddAirplane(ctx context.Context, airplane aviation.Airplane) (*aviation.Airplane, error) {
	if _, ok := r.aa[strings.ToLower(airplane.Registration)]; ok {
		return nil, errors.Errorf("airplane already exists")
	}
	r.aa[strings.ToLower(airplane.Registration)] = &airplane
	return &airplane, nil
}

func getDefaultAirplanes() map[string]*aviation.Airplane {
	return map[string]*aviation.Airplane{
		"g-cnab": {
			Make:         "Jabiru",
			Model:        "UL-450",
			Owner:        "GCNAB Group",
			Registration: "G-CNAB",
			YearBuilt:    2001,
		},
		"g-ufci": {
			Make:         "Cessna",
			Model:        "172",
			Owner:        "Ulster Flying Club",
			Registration: "G-UFCI",
			YearBuilt:    2007,
		},
		"g-cgfz": {
			Make:         "Thruster Air Services",
			Model:        "T600N",
			YearBuilt:    2011,
			Registration: "G-CGFZ",
			Owner:        "NiMicrolights",
		},
	}
}
