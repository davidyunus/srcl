package sailboat

import "github.com/srcl/shipping-yard/domain"

type sailBoat struct {
}

func NewSailBoat() domain.Shipping {
	return &sailBoat{}
}

func (s *sailBoat) StartEngine() (*domain.Ship, error) {
	// do something
	return nil, nil
}

func (s *sailBoat) SpeedUp() error {
	// do something
	return nil
}

func (s *sailBoat) Break() error {
	// do something
	return nil
}
