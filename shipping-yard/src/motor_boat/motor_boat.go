package motorboat

import "github.com/srcl/shipping-yard/domain"

type motorBoat struct {
}

func NewMotorBoat() domain.Shipping {
	return &motorBoat{}
}

func (m *motorBoat) StartEngine() (*domain.Ship, error) {
	// do something
	return nil, nil
}

func (m *motorBoat) SpeedUp() error {
	// do something
	return nil
}

func (m *motorBoat) Break() error {
	// do something
	return nil
}
