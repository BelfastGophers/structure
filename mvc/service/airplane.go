package service

import (
	"context"
	"strings"

	"github.com/belfastgophers/structure/mvc"
	"github.com/pkg/errors"
)

type airplaneService struct {
	repo aviation.AirplaneRepository
}

func NewAirplaneService(repo aviation.AirplaneRepository) *airplaneService {
	return &airplaneService{
		repo: repo,
	}
}

// Airplane will return an airplane by registration if found.
// reg cannot be an empty string or > 6 chars.
func (s *airplaneService) Airplane(ctx context.Context, reg string) (*aviation.Airplane, error) {
	if err := s.validateRegistration(reg); err != nil {
		return nil, errors.WithStack(err)
	}
	airplane, err := s.repo.Airplane(ctx, reg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return airplane, nil
}

// Airplanes will return all stored airplanes.
func (s *airplaneService) Airplanes(ctx context.Context) ([]aviation.Airplane, error) {
	airplanes, err := s.repo.Airplanes(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return airplanes, nil
}

// AddAirplane will add a new airplane to the data store.
// airplane.Registration cannot be empty of > 6 chars.
func (s *airplaneService) AddAirplane(ctx context.Context, airplane aviation.Airplane) (*aviation.Airplane, error) {
	if err := s.validateRegistration(airplane.Registration); err != nil {
		return nil, errors.WithStack(err)
	}
	retval, err := s.repo.AddAirplane(ctx, airplane)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return retval, nil
}

func (s *airplaneService) validateRegistration(reg string) error {
	if strings.TrimSpace(reg) == "" {
		return errors.Errorf("empty registration submitted")
	}
	if len(strings.TrimSpace(reg)) > 6 {
		return errors.Errorf("invalid registration")
	}
	return nil
}
