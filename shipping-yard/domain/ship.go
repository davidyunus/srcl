package domain

type Shipping interface {
	StartEngine() (*Ship, error)
	SpeedUp() error
	Break() error
}

type Ship struct {
	Name     string  `json:"name"`
	Capacity int     `json:"capacity"`
	MaxSpeed int     `json:"maxSpeed"`
	Price    float64 `json:"price"`
}
